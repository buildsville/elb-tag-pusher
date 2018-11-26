Sample deployments
=====

Sample deployments of how to use `elb-tag-pusher`.  
Deploy prometheus and some exporters, publish prometheus with `type:LoadBalancer` service, and monitoring its ELB metrics.

### set aws policy

for elb-tag-pusher

```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "elasticloadbalancing:DescribeLoadBalancers",
                "elasticloadbalancing:DescribeTags"
            ],
            "Resource": "*"
        }
    ]
}
```

for cloudwatch-exporter

```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "cloudwatch:ListMetrics",
                "cloudwatch:GetMetricStatistics"
            ],
            "Resource": "*"
        }
    ]
}
```

### modify yaml files

Set aws authentication in `spec.template.spec.containers[].env` of `elb-tag-pusher.yaml` and `cloudwatch-exporter.yaml`.

```
env:
  - name: AWS_REGION
    value: YOUR_REGION
  - name: AWS_ACCESS_KEY_ID
    value: YOUR_AWS_ACCESS_KEY_ID
  - name: AWS_SECRET_ACCESS_KEY
    value: YOUR_AWS_SECRET_ACCESS_KEY
```

### deploy

Exec command
```
$ kubectl apply -f ./sample/
```

### get metrics

First, get external address.

Exec command
```
$ kubectl get service prometheus-svc
```

and access `http://[EXTERNAL-IP]:[PORT]/status` in your browser.  
You should be able to access the status page of prometheus.  
Please reload several times to make metrics on cloudwatch.  
(If there is no access to the load balancer, cloudwatch-exporter can not create metrics.)  

Wait about 1 minute and access `http://[EXTERNAL-IP]:[PORT]/graph` and enter PromQL,  
and should be able to get the expected metrics.

<img src="https://github.com/buildsville/elb-tag-pusher/blob/master/sample/screenshot.png" alt="screenshot" title="prometheus">
