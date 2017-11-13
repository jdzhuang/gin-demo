[![Go Report Card](https://goreportcard.com/badge/github.com/jdzhuang/gin-demo)](https://goreportcard.com/report/github.com/jdzhuang/gin-demo)

## demo-code depends on gin via docker.

*to build*
```sh
sudo docker build ./ --tag "gin-demo:0.0.1"
```

*to run*
```sh
sudo docker run --rm -d -p 127.0.0.1:80:8080 -p 127.0.0.1:81:8081 "gin-demo:0.0.1"
```

*batch-request*
```sh
for i in $(seq 0 500); do curl "http://127.0.0.1/ping?type=yaml"; usleep 10; done
```

*check-profile*
```sh
curl "http://127.0.0.1:81/profile?type=yaml"
```



