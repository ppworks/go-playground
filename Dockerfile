FROM golang:latest
WORKDIR /go/src/app

RUN go get github.com/pilu/fresh
CMD ["fresh"]
