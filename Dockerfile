FROM golang:1.21-alpine3.18
WORKDIR /app

RUN apk update && apk add --no-cache bash

COPY . .

RUN chmod +x /app/scripts/run.sh
CMD ["/bin/bash", "/app/scripts/run.sh", ">", "docker.log", "2>&1"]