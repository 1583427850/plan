# 使用官方 Golang 镜像作为基础镜像
FROM golang:latest

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件以下载依赖
COPY go.mod .
COPY go.sum .

# 下载依赖
RUN go mod download

# 将当前目录下的所有内容复制到工作目录
COPY . .

# 构建应用
RUN go build -o main .

# 暴露端口
EXPOSE 8080

# 运行应用
CMD ["./main"]

