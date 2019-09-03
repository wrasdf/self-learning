FROM golang:1.12.9-alpine3.10

ENV GO111MODULE=on
WORKDIR /go/src/api/
RUN apk add --no-cache ca-certificates git bash curl
COPY go.mod go.sum ./
RUN go mod download

COPY . ./
