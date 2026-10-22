package main

import (
	"log"

	pb "github.com/augustersej123/ChitChat/grpc"
)

// sendJoin ticks the clock and sends a Join event
func sendJoin(stream pb.ChitChat_ChatClient, clock *LogicalClock, name string) error {
	t := clock.Tick()
	log.Printf("[Client] event=send-join name=%s lamport=%d", name, t)
	return stream.Send(&pb.ClientEvent{
		Timestamp: t,
		Event:     &pb.ClientEvent_Join{Join: &pb.Join{Participant: name}},
	})
}
