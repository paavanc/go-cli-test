
FROM golang:1.12.0-alpine3.9 AS builder
RUN go version
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN apk add build-base
RUN mkdir -p /go/src/go-cli-starter
WORKDIR /go/src/go-cli-starter
COPY . .

RUN set -x && \
    go get github.com/golang/dep/cmd/dep && \
    dep ensure -v



RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -o app .



FROM alpine:3.9
RUN apk --no-cache add ca-certificates


WORKDIR /root/
COPY --from=builder /go/src/go-cli-starter/app .
RUN apk --no-cache add tzdata
RUN chmod +x app

