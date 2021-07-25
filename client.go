package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"github.com/Paulo-Lopes-Estevao/basic-grpc/chat"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000",grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect %v", err)
	}
	defer conn.Close()

	c:= chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "HELLO WORD GRPC",
	}

	response, err := c.SayHello(context.Background(), &message)

	if err != nil {
		log.Fatalf("Error when calling SayHello %v",err)
	}

	log.Printf("Response from Server %s",response.Body)

}
