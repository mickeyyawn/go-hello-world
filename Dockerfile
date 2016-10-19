FROM alpine:3.3

MAINTAINER Mickey Yawn <mickey.yawn@turner.com>

ADD bin/go-hello-world /go-hello-world

ENTRYPOINT ["/go-hello-world"]

EXPOSE 8080
