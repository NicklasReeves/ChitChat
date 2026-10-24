package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"

	pb "github.com/augustersej123/ChitChat/grpc"
)

// readInput sends each line the user types. Blocks until the user leaves
func readInput(stream pb.ChitChat_ChatClient, clock *LogicalClock) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}
		if text == "/leave" {
			break
		}
		if utf8.RuneCountInString(text) > 128 {
			fmt.Println("Message too long (max 128 characters)")
			continue
		}
		if err := sendPublish(stream, clock, text); err != nil {
			log.Printf("event=send-error err=%v", err)
			return
		}
	}

	// Reached on /leave or when stdin closes (Ctrl+D)
	if err := sendLeave(stream, clock); err != nil {
		log.Printf("event=send-error err=%v", err)
	}
	stream.CloseSend()
}