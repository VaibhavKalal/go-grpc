package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/vaibhav/go-grpc/proto"
)

func callSayHelloBidirectionalStreaming(client pb.GreetServiceClient, names *pb.NameList) {
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("failed to establish stream %v", err)
	}

	log.Println("request stream started")
	for _, name := range names.Names {
		if err := stream.Send(&pb.HelloRequest{Name: name}); err != nil {
			log.Fatalf("failed to stream request %v", err)
		}
		time.Sleep(2 * time.Second)
	}

	if err := stream.CloseSend(); err != nil {
		log.Fatalf("failed to close send stream %v", err)
	}

	log.Println("request stream finished")

	log.Println("response stream started")
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to stream respose %v", err)
		}
		log.Println(res.Message)
	}
	log.Println("response stream finished")
}
