package main

import (
	"io"
	"log"

	pb "github.com/vaibhav/go-grpc/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Messages: messages})
		}

		log.Println("message received :", req.Name)
		messages = append(messages, req.Name)
	}
}
