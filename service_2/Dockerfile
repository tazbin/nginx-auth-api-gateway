# server 2

FROM golang:latest

WORKDIR /app

COPY . /app

RUN go build -o main

EXPOSE 3002

ENTRYPOINT [ "/app/main" ]