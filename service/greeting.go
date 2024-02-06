package service

import (
	"context"
	"github.com/lixiang4u/learn-grpc/rpc/greeter"
	"log"
)

type GreetingService struct {
	greeter.UnimplementedGreetingServiceServer
}

func (x GreetingService) SayHello(ctx context.Context, req *greeter.HelloRequest) (*greeter.HelloResponse, error) {
	log.Println("[GreetingService.req]", req.String())
	return &greeter.HelloResponse{Reply: "hello " + req.Greeting}, nil
}

func (x GreetingService) SayHelloAgain(ctx context.Context, req *greeter.HelloRequest) (*greeter.HelloResponse, error) {
	log.Println("[GreetingService.req]", req.String())
	return &greeter.HelloResponse{Reply: "hello " + req.Greeting}, nil
}
