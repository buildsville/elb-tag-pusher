FROM golang:alpine AS builder

RUN apk update && \
    apk add git build-base && \
    rm -rf /var/cache/apk/* && \
    mkdir -p "$GOPATH/src/github.com/buildsville/" && \
    git clone https://github.com/buildsville/elb-tag-exporter.git && \
    mv elb-tag-exporter "$GOPATH/src/github.com/buildsville/" && \
    cd "$GOPATH/src/github.com/buildsville/elb-tag-exporter" && \
    GOOS=linux GOARCH=amd64 go build -o /elb-tag-exporter

FROM alpine:3.7

RUN apk add --update ca-certificates

COPY --from=builder /elb-tag-exporter /elb-tag-exporter

ENTRYPOINT ["/elb-tag-exporter"]
