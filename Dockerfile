FROM golang:1.8.5-alpine3.6 AS build-gin-demo
MAINTAINER  jdzhuang<jiadong.zhuang@qq.com> 

RUN mkdir -p /go/src/gin-demo/build
WORKDIR /go/src/gin-demo
COPY . /go/src/gin-demo/
RUN go list ./...|grep -v vendor|xargs go test
RUN go build -v -o build/gin-demo.bin

FROM alpine:3.6
EXPOSE 8080 8081
COPY --from=build-gin-demo /go/src/gin-demo/build/gin-demo.bin /usr/local/bin/
ENTRYPOINT ["gin-demo.bin", "", ""]

