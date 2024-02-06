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

	var client = greeter.NewHelloServiceClient(conn)
	resp, err := client.SayHello(ctx, &greeter.HelloRequest{Greeting: time.Now().Format("2006-01-02 15:04:05")})
	if err != nil {
		log.Println("[client.SayHello]", err.Error())
		return
	}
	log.Println("[resp]", resp.Reply)
	log.Println("[resp]", resp.String())
}
