FROM golang:1.18-alpine

# default APP_PORT set 8080, can change on create container
ENV APP_PORT=8080

RUN mkdir app
COPY main.go app

EXPOSE ${APP_PORT}

CMD go run app/main.go