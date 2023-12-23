# 基础镜像
# docker run -itd name=test golang:1.21.5-alpine3.18 ash
FROM golang:1.21.5-alpine3.18
# go env
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
# 备份原始安装源
RUN cp /etc/apk/repositories /etc/apk/repositories.bak
# 修改为国内源
#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
# RUN sed -i 's/https:\/\/dl-cdn.alpinelinux.org/http:\/\/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk update
RUN apk upgrade
# 安装基础软件
# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o translate main.go
RUN apk add translate-shell
RUN apk add sqlite-dev
RUN apk add bash
RUN apk add git
RUN apk add build-base
RUN git clone -b docker https://github.com/zhangyiming748/freeTranslate.git /go/freeTranslate
WORKDIR /go/freeTranslate
RUN go mod tidy
RUN go mod vendor
RUN go build -o trans main.go
# 准备软件
CMD ["/go/freeTranslate/trans"]
#docker build -t trans:v1 .
#docker run -idt --name=trans -v /d/srt:/srt -e APPID={your baidu appid} -e KEY={your baidu key} trans:v1 ash
