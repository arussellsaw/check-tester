FROM alpine:3.2

ADD ./check-tester /tmp/check-tester

EXPOSE 8888

ENTRYPOINT ["/tmp/check-tester"]
