FROM golang:alpine

ENV GO111MODULE=on

RUN apk update && \
    apk upgrade && \
    apk add --no-cache \
    bash \
    git \
    gcc \
    libc-dev
RUN go get github.com/pilu/fresh
RUN go get github.com/pressly/goose/cmd/goose

WORKDIR /go/src/github.com/ppworks/go-playground

CMD ["fresh"]
