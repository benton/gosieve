FROM busybox:latest

ADD pkg/gosieve /gosieve

ENTRYPOINT ["/gosieve"]
