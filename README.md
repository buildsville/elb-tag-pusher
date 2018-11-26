elb-tag-pusher
=====
`elb-tag-pusher` pushes tags information of `aws elastic load balancer` to pushGateWay.  
This is used in prometheus to relate cloudwatch metrics of elb with the service of kubernetes.  
For details to use, please see `sample/README.md`.

Can see command line flags

```
./elb-tag-pusher -h
```

### Caution

If the key of the tag contains a symbol, it will be replaced with an underscore.  
If the value of the tag contains `/`, it will be replaced with an underscore.  
These are due to specification limitations of label of prometheus.  

### Why use pushgateway without creating exporter

It was difficult for me to create an exporter that gives a dynamic label using `prometheus/client_golang`...

### docker image  
https://hub.docker.com/r/masahata/elb-tag-pusher/
