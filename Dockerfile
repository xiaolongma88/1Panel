# 使用官方的 Ubuntu 镜像作为基础镜像
FROM ubuntu:22.04

# 安装 Go
RUN apt-get update && apt-get install -y golang

# 设置工作目录
WORKDIR /app

# 复制当前目录的内容到工作目录
COPY . .

# 编译 Go 程序
CMD ["cd ./cmd/server"]
RUN go build -o myapp .

# 运行 Go 程序
CMD ["./myapp"]
