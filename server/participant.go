package main

import pb "github.com/augustersej123/ChitChat/grpc"

// Participant is one connected client. Out is drained by that
// client's sender goroutine in the chat handler
type Participant struct {
	ID  string
	Out chan *pb.ChatMessage
}

func newParticipant(id string) *Participant {
	return &Participant{
		ID:  id,
		Out: make(chan *pb.ChatMessage, 32), // buffered so a slow client won't block the whole server
	}
}
