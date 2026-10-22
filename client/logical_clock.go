package main

import "sync"

// Same code as the server/logical_clock.go
// LogicalClock is a Lamport clock.
type LogicalClock struct {
	mu   sync.Mutex // protects time
	time int64
}

// Tick is for local events: time = time + 1
func (c *LogicalClock) Tick() int64 { // Tick is for local events, e.g. before the client sends a join/publish/leave.
	c.mu.Lock()
	defer c.mu.Unlock()
	c.time++
	return c.time
}

// Update is for receiving events: time = max(local, received) + 1
func (c *LogicalClock) Update(received int64) int64 { //Used when a message is received from another participant.
	c.mu.Lock()
	defer c.mu.Unlock()
	if received > c.time {
		c.time = received
	}
	c.time++
	return c.time
}

// Now reads the time without changing it.
func (c *LogicalClock) Now() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.time
}
