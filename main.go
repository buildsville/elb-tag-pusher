package main

import (
	"flag"
	"regexp"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/prometheus/common/log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elb"
)

const (
	defaultPushIntervalSec = 60
	defaultPushGateWayAddr = "http://localhost:9091"
	defaultJobName         = "push_elb_tag"
	defaultMetricsName     = "aws_elb_tags"
	defalutELBNameLabelKey = "load_balancer_name"
	defaultReplaceChar     = "_"
)

var pushAddr = flag.String("pushGateWayAddr", defaultPushGateWayAddr, "push metrics gateway address.")
var interval = flag.Int("pushInterval", defaultPushIntervalSec, "Interval to push metrics.")
var jobName = flag.String("jobName", defaultJobName, "job name of metrics.")
var metricsName = flag.String("metricsName", defaultMetricsName, "metrics name of elb tag info.")
var elbNameLabelKey = flag.String("elbNameLabelKey", defalutELBNameLabelKey, "the key of elb name label.")
var replace = flag.String("replaceChar", defaultReplaceChar, "a character to replace unusable characters in labels.")

var elbSession = func() *elb.ELB {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	return elb.New(sess)
}()

func getTagDescriptions() ([]*elb.TagDescription, error) {
	var tags []*elb.TagDescription
	input := &elb.DescribeLoadBalancersInput{}
	ret, err := elbSession.DescribeLoadBalancers(input)
	if err != nil {
		return nil, err
	}
	for _, lb := range ret.LoadBalancerDescriptions {
		taginput := &elb.DescribeTagsInput{
			LoadBalancerNames: []*string{lb.LoadBalancerName},
		}
		tag, err := elbSession.DescribeTags(taginput)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag.TagDescriptions...)
	}
	return tags, nil
}

func main() {
	flag.Parse()
	keyReg := regexp.MustCompile(`[ -/:-@{-~]`)
	valReg := regexp.MustCompile(`/`)
	for {
		tagDescs, err := getTagDescriptions()
		if err != nil {
			log.Errorln(err)
		}
		awsElbTags := prometheus.NewGauge(prometheus.GaugeOpts{
			Name: *metricsName,
			Help: "tag key and value on aws elastic load balancer.",
		})
		awsElbTags.Set(float64(1))
		for _, tagDesc := range tagDescs {
			push := push.New(*pushAddr, *jobName).Collector(awsElbTags)
			push.Grouping(*elbNameLabelKey, *tagDesc.LoadBalancerName)
			for _, t := range tagDesc.Tags {
				push.Grouping(keyReg.ReplaceAllString(*t.Key, *replace), valReg.ReplaceAllString(*t.Value, *replace))
			}
			if err := push.Push(); err != nil {
				log.Errorf("Could not push completion time to Pushgateway:", err)
			}
		}
		time.Sleep(time.Duration(*interval) * time.Second)
	}
}
