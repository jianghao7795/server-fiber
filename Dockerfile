FROM golang:1.22.8-alpine AS builder

LABEL org.opencontainers.image.authors="jianghao"

ENV GOPROXY=https://goproxy.cn,direct
# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
ENV GO111MODULE=on

WORKDIR /app
COPY . /app
RUN go build -o fiber cmd/main.go

FROM rockylinux:9-minimal AS runner
WORKDIR /app

COPY --from=builder /app/fiber .
COPY --from=builder /app/config.yaml ./conf/config.yaml
COPY --from=builder /app/rbac_model.conf .
COPY --from=builder /app/docs/ ./docs/

EXPOSE 3100
CMD ["/app/fiber", "-c", "./conf/"]
# ENTRYPOINT ["/app/server-fiber", "-c", "config.yaml"]
