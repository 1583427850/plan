# 使用官方 Golang 镜像作为基础镜像
FROM golang:latest

# 设置工作目录
WORKDIR /app


# 将当前目录下的所有内容复制到工作目录
COPY . .

RUN go mod tidy

# 构建应用
RUN go build -o main main.go

# 暴露端口
EXPOSE 8080

# 运行应用
CMD ["./main"]
