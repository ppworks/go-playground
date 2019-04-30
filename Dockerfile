FROM golang:alpine

RUN apk update && \
    apk upgrade && \
    apk add --no-cache \
    bash \
    git

WORKDIR /go/src/app
RUN go get github.com/pilu/fresh
CMD ["fresh"]
