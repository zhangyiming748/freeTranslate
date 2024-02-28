# 基础镜像
# docker run -itd name=test golang:1.21.5-alpine3.18 ash
FROM golang:1.22.0-alpine3.19
# go env
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GOBIN=/go/bin
# 备份原始安装源
RUN cp /etc/apk/repositories /etc/apk/repositories.bak
# 修改为国内源
RUN sed -i 's/https:\/\/dl-cdn.alpinelinux.org/http:\/\/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
RUN apk update
RUN apk upgrade
# 安装基础软件
# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o translate main.go
RUN apk add translate-shell sqlite-dev bash git build-base
RUN git clone https://github.com/zhangyiming748/freeTranslate.git /root/freeTranslate
WORKDIR /root/freeTranslate
RUN go get -u
RUN go mod tidy
RUN go mod vendor
# 准备软件
#CMD ["go","run","/root/freeTranslate/main.go"]
#docker build -t trans:v1 .
#docker run --name=test -it --rm trans:v1 ash
