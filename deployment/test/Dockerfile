FROM golang:1.12.1-alpine

ENV GO111MODULE on

RUN apk update
RUN apk add curl git 
#gcc musl-dev bash

LABEL maintainer="Fajrul Aulia <auliafajrul7@gmail.com>"

RUN mkdir /app
WORKDIR /app
COPY . .
RUN go build -o main cmd/main.go
EXPOSE 8084
CMD ["/app/main"]