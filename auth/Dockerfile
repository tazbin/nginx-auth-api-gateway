# server auth

FROM golang:latest

WORKDIR /app

COPY . /app

RUN go build -o main

EXPOSE 3000

ENTRYPOINT [ "/app/main" ]