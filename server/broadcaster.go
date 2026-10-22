package main

import (
	"fmt"
	"log"

	pb "github.com/augustersej123/ChitChat/grpc"
)

type EventType int

const (
	EventJoin EventType = iota
	EventLeave
	EventPublish
)

// Event is what the Chat handlers send to the broadcaster
type Event struct {
	Type        EventType
	Participant *Participant //used by join and leave
	Content     string       //used by publish
	ClientTime  int64        //the clients's Lamport time(0 if none)
}

type Broadcaster struct {
	events       chan Event
	participants map[string]*Participant
	clock        *LogicalClock
}

func NewBroadcaster(clock *LogicalClock) *Broadcaster {
	return &Broadcaster{
		events:       make(chan Event),
		participants: make(map[string]*Participant),
		clock:        clock,
	}
}

// Send is how handlers submit events
func (b *Broadcaster) Send(e Event) {
	b.events <- e
}

// Run must be started with "go b.Run()"
func (b *Broadcaster) Run() {
	for e := range b.events {
		switch e.Type {
		case EventJoin:
			b.participants[e.Participant.ID] = e.Participant
			t := b.clock.Update(e.ClientTime)
			text := fmt.Sprintf("Participant %s joined the chat at logical time %d", e.Participant.ID, t)
			log.Printf("[Server] event=broadcast-join client=%s lamport=%d", e.Participant.ID, t)
			b.deliver("Server", text, t)

		case EventLeave:
			p, ok := b.participants[e.Participant.ID]
			if !ok {
				continue //already removed
			}
			delete(b.participants, e.Participant.ID)
			close(p.Out)
			t := b.clock.Update(e.ClientTime)
			text := fmt.Sprintf("Participant %s left the chat at logical time %d", e.Participant.ID, t)
			log.Printf("[Server] event=broadcast-leave client=%s lamport=%d", e.Participant.ID, t)
			b.deliver("Server", text, t)

		case EventPublish:
			t := b.clock.Update(e.ClientTime)
			log.Printf("[Server] event=broadcast-publish client=%s lamport=%d", e.Participant.ID, t)
			b.deliver(e.Participant.ID, e.Content, t)
		}
	}
}

func (b *Broadcaster) deliver(sender, content string, t int64) {
	msg := &pb.ChatMessage{Sender: sender, Content: content, Timestamp: t}
	for id, p := range b.participants {
		select {
		case p.Out <- msg:
			log.Printf("[Server] event=deliver client=%s lamport=%d", id, t)
		default:
			log.Printf("[Server] event=deliver client=%s lamport=%d dropped message", id, t)
		}
	}
}
