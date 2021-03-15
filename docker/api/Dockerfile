FROM golang:latest

COPY ./api /go/src/api

WORKDIR /go/src/api

ENV GO111MODULE=on

RUN go mod download

RUN go get -u github.com/cosmtrek/air