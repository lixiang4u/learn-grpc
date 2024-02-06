## 1、定义proto协议
> pb/HelloService.proto


## 2、编译proto文件
> protoc  --go_out=. --go-grpc_out=.  .\pb\HelloService.proto

## 3、编写服务端和客户端
> [server.go](server.go)
> 
> [client.go](client.go)


## 4、启动服务端和客户端验证
> go get 
> 
> go run server.go
> 
> go run client.go