package main

import (
	"context"
	pb "golang-book/chapter2/daemon-logging/logservice"
	"log"
	"log/syslog"
	"net"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type logServer struct {
	pb.UnimplementedLogServiceServer
}

func (s *logServer) LogMessage(ctx context.Context, req *pb.LogRequest) (*pb.LogResponse, error) {
	processName := ExtractProcessName(req.ProcessName)
	logMessage := FormatLogMessage(req.Message)
	// Send log entry to syslog
	syslogWriter, err := syslog.New(syslog.LOG_INFO, processName)
	if err == nil {
		syslogWriter.Info(logMessage)
	}

	logrus.WithFields(logrus.Fields{
		"timestamp": time.Now().Format(time.RFC3339),
	}).Info(req.Message)

	return &pb.LogResponse{Success: true}, nil
}

func ExtractProcessName(fullPath string) string {
	splitPath := strings.Split(fullPath, "/")
	return splitPath[len(splitPath)-1]
}

func FormatLogMessage(message string) string {
	return strings.TrimSpace(message)
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLogServiceServer(s, &logServer{})
	log.Println("gRPC Logging Server is running on port 50051...")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
