package main

import (
	"context"
	"github.com/lixiang4u/learn-grpc/rpc/greeter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:3030", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("[grpc.Dial]", err.Error())
		return
	}
	defer func() { _ = conn.Close() }()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var greetingClient = greeter.NewGreetingServiceClient(conn)
	resp, err := greetingClient.SayHelloAgain(ctx, &greeter.HelloRequest{Greeting: time.Now().Format("2006-01-02 15:04:05")})
	if err != nil {
		log.Println("[greetingClient.SayHello]", err.Error())
		return
	}
	log.Println("[greetingClient.resp]", resp.Reply)
	log.Println("[greetingClient.resp]", resp.String())

	var helloClient = greeter.NewHelloServiceClient(conn)
	resp, err = helloClient.SayHello(ctx, &greeter.HelloRequest{Greeting: time.Now().Format("2006-01-02 15:04:05")})
	if err != nil {
		log.Println("[helloClient.SayHello]", err.Error())
		return
	}
	log.Println("[helloClient.resp]", resp.Reply)
	log.Println("[helloClient.resp]", resp.String())
}
