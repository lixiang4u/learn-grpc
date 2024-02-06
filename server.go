package main

import (
	"github.com/lixiang4u/learn-grpc/rpc/greeter"
	"github.com/lixiang4u/learn-grpc/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	listen, err := net.Listen("tcp", ":3030")
	if err != nil {
		log.Println("[net.Listen]", err.Error())
		return
	}

	var s = grpc.NewServer()
	greeter.RegisterHelloServiceServer(s, &service.HelloService{})
	greeter.RegisterGreetingServiceServer(s, &service.GreetingService{})
	if err := s.Serve(listen); err != nil {
		log.Println("[s.Serve]", err.Error())
	}
}
