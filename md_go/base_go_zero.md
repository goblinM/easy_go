# go_zero
#### 基本入门
```
    安装go-zero:  GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/tal-tech/go-zero
    安装goctl:  GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/tal-tech/go-zero/tools/goctl
    快速生成api:
         goctl api new greet
         cd greet
         go mod init
         go mod tidy  // 生成go.mod
         go run greet.go -f etc/greet-api.yaml  // 启动服务
    
```