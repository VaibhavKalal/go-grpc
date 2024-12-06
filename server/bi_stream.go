package main

import (
	"io"
	"log"
	"time"

	pb "github.com/vaibhav/go-grpc/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
	log.Println("request streaming started")
	var names []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed while streaming request %v", err)
		}
		log.Println("Name received :", req.Name)
		names = append(names, req.Name)
	}
	log.Println("request streaming finished")

	log.Println("response streaming started")
	for _, name := range names {
		if err := stream.Send(&pb.HelloResponse{Message: "Hello " + name}); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	log.Println("response streaming finished")

	return nil
}
