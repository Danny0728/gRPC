package main

import (
	"Chat/chat"
	"bufio"
	"context"
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	c := chat.NewChatServiceClient(conn)

	dmsReader := bufio.NewReader(os.Stdin)
	body, _ := dmsReader.ReadString('\n')

	message := &chat.Message{
		Body: body,
	}

	response, err := c.SayHello(context.Background(), message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from the Server: %s", response.Body)
}
