docker build -t twilio . &&
docker run --env-file ./.env --publish 6060:8080 --name test --rm twilio
