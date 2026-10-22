package main

import (
	"io"
	"log"
	"os"

	pb "github.com/augustersej123/ChitChat/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// log handler goes into main.
	f, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	// os.O_CREATE If file doesnt exist we create it
	//os.O_Wronly Opens file for write only
	//os.O_Append makes every write go to the end of the document, so new content doesnt overwrite from the top
	if err != nil {
		log.Fatalf("could not open the log file:%v", err)
	}

	defer f.Close()
	log.SetOutput(io.MultiWriter(os.Stdout, f))     //multiwriter, since we need logs in both appendix and repo
	log.SetFlags(log.LstdFlags | log.Lmicroseconds) //Timestamp format to microseconds

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("[Client] failed to connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewChitChatClient(conn)
	_ = client
	log.Println("[Client] [Connect] connected to server")
}
