package main

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"strings"

	pb "golang-book/chapter2/daemon-logging/logservice"

	"google.golang.org/grpc"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <message>", os.Args[0])
	}

	message := strings.Join(os.Args[1:], " ") // Take log message as a whole string

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewLogServiceClient(conn)
	processName := filepath.Base(os.Args[0])

	_, err = client.LogMessage(context.Background(), &pb.LogRequest{
		ProcessName: processName,
		Message:     message,
	})
	if err != nil {
		log.Fatalf("Failed to send log: %v", err)
	}
}
