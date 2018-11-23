FROM golang:alpine AS builder

RUN apk update && \
    apk add git build-base && \
    rm -rf /var/cache/apk/* && \
    mkdir -p "$GOPATH/src/github.com/buildsville/" && \
    git clone https://github.com/buildsville/elb-tag-pusher.git && \
    mv elb-tag-pusher "$GOPATH/src/github.com/buildsville/" && \
    cd "$GOPATH/src/github.com/buildsville/elb-tag-pusher" && \
    GOOS=linux GOARCH=amd64 go build -o /elb-tag-pusher

FROM alpine:3.7

RUN apk add --update ca-certificates

COPY --from=builder /elb-tag-pusher /elb-tag-pusher

ENTRYPOINT ["/elb-tag-pusher"]
