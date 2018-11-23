elb-tag-pusher
=====
`elb-tag-pusher` pushes aws elastic load balancer's tag information to pushGateWay

can see command line flags

```
./elb-tag-pusher -h
```

## aws role

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

## caution

If the key of the tag contains a symbol, it will be replaced with an underscore  
If the value of the tag contains `/`, it will be replaced with an underscore  
These are due to specification limitations of label of prometheus  

docker image  
https://hub.docker.com/r/masahata/elb-tag-pusher/
