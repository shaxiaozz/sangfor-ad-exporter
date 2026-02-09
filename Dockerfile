FROM golang:1.25.4-alpine as builder
WORKDIR /data/sangfor-ad-exporter-code
ENV GOPROXY=https://goproxy.cn
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add --no-cache upx ca-certificates tzdata
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o sangfor-ad-exporter

FROM golang:1.25.4-alpine as runner
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /data/sangfor-ad-exporter-code/sangfor-ad-exporter /sangfor-ad-exporter
RUN mkdir /etc/sangfor-ad-exporter
EXPOSE 9098
CMD ["/sangfor-ad-exporter","start"]