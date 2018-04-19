FROM scratch

COPY ./twilio-api /

CMD ["/twilio-api"]
EXPOSE 8080
