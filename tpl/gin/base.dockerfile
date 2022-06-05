FROM golang:alpine as base

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk add tzdata
RUN apk add make git

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /build

COPY go.mod go.mod
COPY go.sum go.sum
RUN  go mod download && \
    go install github.com/go-delve/delve/cmd/dlv@latest

CMD sh
