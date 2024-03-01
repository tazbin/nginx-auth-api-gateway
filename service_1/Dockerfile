# server 1

FROM golang:latest

WORKDIR /app

COPY . /app

RUN go build -o main

EXPOSE 3001

ENTRYPOINT [ "/app/main" ]