package main

import (
	"io"
	"log"
	"net"
	"os"

	pb "github.com/augustersej123/ChitChat/grpc"
	"google.golang.org/grpc"
)

func main() {
	f, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("could not open the log file: %v", err)
	}
	defer f.Close()
	log.SetOutput(io.MultiWriter(os.Stdout, f)) // terminal + file
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("[Server] failed to listen: %v", err)
	}

	clock := &LogicalClock{}
	b := NewBroadcaster(clock)
	go b.Run()

	s := grpc.NewServer()
	pb.RegisterChitChatServer(s, &server{broadcaster: b})
	log.Println("[Server] [Startup] listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("[Server] failed to serve: %v", err)
	}
}
