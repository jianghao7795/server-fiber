FROM golang:1.22.8-alpine AS builder

LABEL org.opencontainers.image.authors="jianghao"

ENV GOPROXY=https://goproxy.cn,direct
# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
ENV GO111MODULE=on

WORKDIR /app
COPY . /app
RUN go build -o server-fiber cmd/main.go

FROM rockylinux:9 AS runner
WORKDIR /app

COPY --from=builder /app/server-fiber .
COPY --from=builder /app/conf/ ./conf/
COPY --from=builder /app/rbac_model.conf .
COPY --from=builder /app/docs/ ./docs/

EXPOSE 3100
CMD ["/app/server-fiber", "-c", "./conf/"]
# ENTRYPOINT ["/app/server-fiber", "-c", "config.yaml"]
