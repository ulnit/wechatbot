#64bit
###
 # @Description: 
 # @Version: 1.0
 # @Autor: Sean
 # @Date: 2023-03-18 21:00:40
 # @LastEditors: Sean
 # @LastEditTime: 2023-07-22 17:35:59
### 
GOOS=windows GOARCH=amd64 go build -o bin/wechatbot-amd64.exe main.go

#32-bit
GOOS=windows GOARCH=386 go build -o bin/wechatbot-386.exe main.go

# 64-bit
GOOS=linux GOARCH=amd64 go build -o bin/wechatbot-amd64-linux main.go
# windows下
# $env:GOOS="linux" 
# $env:GOARCH="amd64" 
# $env:CGO_ENABLED=0 
# go build -o bin/wechatbot-amd64-linux main.go

# 32-bit
GOOS=linux GOARCH=386 go build -o bin/wechatbot-386-linux main.go

# 64-bit 
GOOS=darwin GOARCH=amd64 go build -o bin/wechatbot-amd64-darwin main.go 

# 32-bit 
GOOS=darwin GOARCH=386 go build -o bin/wechatbot-386-darwin main.go
