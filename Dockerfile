FROM golang:1.10-alpine
ADD . /go/src/github.com/mcappleman/twilio-api
RUN go install github.com/mcappleman/twilio-api
ENTRYPOINT /go/bin/twilio-api
EXPOSE 8080
