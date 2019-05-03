FROM golang:alpine

RUN apk update && \
    apk upgrade && \
    apk add --no-cache \
    bash \
    git

RUN go get github.com/pilu/fresh

WORKDIR /go/src/github.com/ppworks/go-playground
CMD ["fresh"]
