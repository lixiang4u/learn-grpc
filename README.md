## 1、定义proto协议
> pb/HelloService.proto


## 2、编译proto文件
> protoc  --go_out=. --go-grpc_out=.  .\pb\HelloService.proto

