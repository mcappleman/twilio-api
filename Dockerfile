FROM golang:1.10
ADD . /go/src/github.com/mcappleman/twilio-api

# Use bash instead of sh
RUN rm /bin/sh && ln -s /bin/bash /bin/sh
RUN cat /go/src/github.com/mcappleman/twilio-api/.env > ~/.bash_profile

RUN go install github.com/mcappleman/twilio-api
ENTRYPOINT /go/bin/twilio-api
EXPOSE 8080

#RUN source ~/.bash_profile && echo $LOG_FILE_PATH
