#  main -m release  线上执行命令


#!/bin/sh
out="bin"
echo "build file to ./$out"

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./$out/api  -tags v1 -v main.go

# 压缩二进制文件
upx ./$out/api

echo "build success."