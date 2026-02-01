# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# 设置 Go 代理为国内镜像，解决网络超时问题
ENV GOPROXY=https://goproxy.cn,direct

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/server

# Run stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8081

CMD ["./main"]
