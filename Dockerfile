FROM golang:1.13.4

MAINTAINER XieWei(1156143589@qq.com)

EXPOSE 8888 8081

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/Shanghai" > /etc/timezone
RUN dpkg-reconfigure -f noninteractive tzdata

WORKDIR /go/src/github.com/wuxiaoxiaoshen/go-anything
RUN echo $PWD
COPY . /go/src/github.com/wuxiaoxiaoshen/go-anything

RUN apt-get update && apt-get install -q -y vim nginx postgresql-client git openssh-client cron && apt-get clean;\
    go mod vendor;\
    make remove;\
    make build;\
    echo Succeed!
CMD ["bash","-c", "go run /go/src/github.com/wuxiaoxiaoshen/go-anything/main.go"]