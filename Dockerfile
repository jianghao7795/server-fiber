FROM golang:1.22.2-alpine as builder

MAINTAINER "jianghao"

ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE=on

WORKDIR /app
COPY . /app
RUN go build -o server-fiber ./main.go

EXPOSE 3100
CMD ["/app/server-fiber", "-c", "config.yaml"]
# ENTRYPOINT ["/app/server-fiber", "-c", "config.yaml"]