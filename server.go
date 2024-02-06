package main

import (
	"context"
	"github.com/lixiang4u/learn-grpc/rpc/greeter"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	greeter.UnimplementedHelloServiceServer
}

func (server) SayHello(ctx context.Context, req *greeter.HelloRequest) (*greeter.HelloResponse, error) {
	log.Println("[request]", req.Greeting)
	var resp = greeter.HelloResponse{}
	resp.Reply = "hello " + req.Greeting
	return &resp, nil
}

func main() {

	listen, err := net.Listen("tcp", ":3030")
	if err != nil {
		log.Println("[net.Listen]", err.Error())
		return
	}

	var s = grpc.NewServer()
	greeter.RegisterHelloServiceServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Println("[s.Serve]", err.Error())
	}
}
