rm twilio-api &&
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo &&
docker build -t twilio . --no-cache &&
docker run --env-file ./.env --publish 127.0.0.1:8090:8080 --name twilio --rm twilio
