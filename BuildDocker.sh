docker build -t twilio . &&
docker run --env-file ./.env --publish 127.0.0.1:8090:8080 --name test --rm twilio
