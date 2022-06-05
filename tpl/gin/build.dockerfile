FROM golang:base as builder

WORKDIR /build
COPY . .

RUN ls /build/ && make release

# release
FROM alpine:latest

ENV VERSION=1.0 \
    GOTRACEBACK=crash

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /build/{{ .AppName }} /app/{{ .AppName }}

COPY --from=builder /go/bin/dlv /usr/bin/dlv
ENV TZ=Asia/Shanghai

# 挂载容器目录
VOLUME ["/app/conf", "/data/corefile" ]

EXPOSE 3000

# ENV TINI_VERSION v0.19.0
# ENTRYPOINT ["tini", "--"]
CMD ["/app/{{ .AppName }}","up","--mode=all","-c=/app/conf/app.ini"]
