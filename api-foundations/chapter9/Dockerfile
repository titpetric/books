FROM alpine:latest

ADD flags /flags
RUN chmod +x /flags ; sync; sleep 1

WORKDIR /

ENV GO_PORT 8080
EXPOSE $GO_PORT

ENTRYPOINT ["/flags"]
