package main

import (
	"Chat/chat"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen of port 9000 %v", err)
	}
	s := chat.Server{}
	grpcNewServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcNewServer, &s)
	if err := grpcNewServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC over port 9000: %v", err)
	}
}
