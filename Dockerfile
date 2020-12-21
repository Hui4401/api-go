FROM golang:1.15

LABEL maintainer="Assassin"

WORKDIR /root/api-go

COPY . .

ENV GO111MODULE=on
ENV GOPROXY="https://mirrors.aliyun.com/goproxy,direct"

CMD go run main.go
