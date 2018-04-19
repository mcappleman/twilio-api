FROM golang:1.10-alpine

ADD ./twilio-api /go/bin

ENTRYPOINT /go/bin/twilio-api
EXPOSE 8080
