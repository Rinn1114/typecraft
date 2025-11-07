# --- 构建阶段 ---
FROM golang:1.21-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件，利用 Docker 的缓存机制
COPY go.mod go.sum ./
# 下载依赖
RUN go mod download

# 复制所有源代码
COPY . .

# 编译 Go 应用为一个静态的、在 Linux 环境下运行的二进制文件
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main .

# --- 运行阶段 ---
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 从构建阶段复制编译好的二进制文件
COPY --from=builder /app/main .

# 声明容器将暴露 8080 端口
EXPOSE 8080

# 定义容器启动时要执行的命令
CMD ["./main"]
