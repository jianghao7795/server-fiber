FROM golang:1.22.4-alpine3.20

MAINTAINER "jianghao"

ENV GOPROXY https://goproxy.cn,direct
# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
ENV GO111MODULE=on

WORKDIR /app
COPY . /app
RUN go build -o server-fiber .

EXPOSE 3100
CMD ["/app/server-fiber", "-c", "config.yaml"]
# ENTRYPOINT ["/app/server-fiber", "-c", "config.yaml"]