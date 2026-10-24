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

// sendPublish ticks the clock and sends a Publish event
func sendPublish(stream pb.ChitChat_ChatClient, clock *LogicalClock, content string) error {
	t := clock.Tick()
	log.Printf("event=send-publish lamport=%d", t)
	return stream.Send(&pb.ClientEvent{
		Timestamp: t,
		Event:     &pb.ClientEvent_Publish{Publish: &pb.Publish{Content: content}},
	})
}

// sendLeave ticks the clock and sends a Leave event
func sendLeave(stream pb.ChitChat_ChatClient, clock *LogicalClock) error {
	t := clock.Tick()
	log.Printf("event=send-leave lamport=%d", t)
	return stream.Send(&pb.ClientEvent{
		Timestamp: t,
		Event:     &pb.ClientEvent_Leave{Leave: &pb.Leave{}},
	})
}
