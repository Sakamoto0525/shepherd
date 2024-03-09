FROM golang:1.22-alpine as local

ENV ROOT=/go/src/github.com/Sakamoto0525/shepherd

WORKDIR ${ROOT}
COPY . .

RUN apk add --no-cache alpine-sdk git && \
    go install github.com/cosmtrek/air@latest
