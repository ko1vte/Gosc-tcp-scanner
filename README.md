# Gosc/Tcp Scanner
Go编写的简单的TCP扫描器

你可以这样使用它
```Bash
go run ./src/main.go -p 1024 -g 100 127.0.0.1
```
p为扫描的端口范围：0-p

g为goroutine的数量：g

后接ipv4的地址。

你也可以将其编译后使用
```Bash
go build ./src/gosc.go
gosc -p 445 127.0.0.1
