FROM golang:1.8.4-alpine3.6
MAINTAINER  jdzhuang<jiadong.zhuang@qq.com> 
RUN mkdir -p /usr/local/go-runner
ADD ./gin-demo.bin /usr/local/go-runner/
WORKDIR /usr/local/go-runner
ENTRYPOINT ["./gin-demo.bin", "", ""]


