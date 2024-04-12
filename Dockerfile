FROM golang:1.22.1-alpine

MAINTAINER "jianghao"

ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE=on

WORKDIR /app
COPY config.yaml /app/config.yaml
ADD . /app
RUN go build -o server-fiber ./main.go

EXPOSE 3100
CMD ["/app/server-fiber", "-c", "config.yaml"]