FROM golang:1.10-alpine

ADD . /go/src/github.com/mcappleman/twilio-api
RUN apk add --no-cache bash
RUN bash -c 'ls -la /go/src/github.com/mcappleman/twilio-api'
RUN apk add --no-cache git \
	&& go get github.com/mcappleman/twilio-api \
	&& apk del git \
	&& go install github.com/mcappleman/twilio-api
RUN bash -c 'ls -la /go/src/github.com/mcappleman/twilio-api'
RUN apk del bash

ENTRYPOINT /go/bin/twilio-api
EXPOSE 8080
