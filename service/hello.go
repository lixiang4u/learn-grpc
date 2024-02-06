package service

import (
	"context"
	"github.com/lixiang4u/learn-grpc/rpc/greeter"
	"log"
)

type HelloService struct {
	greeter.UnimplementedHelloServiceServer
}

func (x HelloService) SayHello(ctx context.Context, req *greeter.HelloRequest) (*greeter.HelloResponse, error) {
	log.Println("[HelloService.req]", req.String())
	return &greeter.HelloResponse{Reply: "hello " + req.Greeting}, nil
}

func (x HelloService) SayHelloAgain(ctx context.Context, req *greeter.HelloRequest) (*greeter.HelloResponse, error) {
	log.Println("[HelloService.req]", req.String())
	return &greeter.HelloResponse{Reply: "hello " + req.Greeting}, nil
}
