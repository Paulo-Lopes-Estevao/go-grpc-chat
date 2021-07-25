package chat

import (
	"golang.org/x/net/context"
	"log"
)

type Server struct {
	UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context,message *Message) (*Message, error){
	log.Println("Received message body from client: %s", message.Body)
	return &Message{
		Body: "Hello From the Server!",
	}, nil
}