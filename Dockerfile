FROM golang:1.10-alpine

COPY ./twilio-api /

CMD ["/twilio-api"]
EXPOSE 8080
