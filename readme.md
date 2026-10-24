# ChitChat

A distributed chat service built with Go and gRPC. Clients join a chat, publish messages (max 128 characters) and leave. The server broadcasts every event to all participants, and every message carries a Lamport timestamp.

## Requirements

- Go 1.27 or newer
- No other setup needed. The generated gRPC code is already in `grpc/`

## Running

All commands are run from the project root (the folder with `go.mod`).

**1. Start the server** (listens on `localhost:50051`):

```bash
go run ./server
```

**2. Start a client** in a new terminal, giving it a name:

```bash
go run ./client alice
```

Start more clients in separate terminals, each with its own name, e.g. `go run ./client bob` and `go run ./client carol`.

## Using the chat

- Type a message and press **Enter** to send it to everyone
- Type `/leave` to leave the chat
- Pressing **Ctrl+C** also leaves. The server detects the dropped connection and broadcasts the leave

Messages are printed as `[lamport time] sender: message`, for example:

```
[4] alice: hello everyone
```

## Logs

Logs are written to the folder you run the commands from:

| File | Contents |
|---|---|
| `server.log` | Every join, publish, leave and delivery on the server (also printed in the server terminal) |
| `client_<name>.log` | That client's sent and received events with Lamport times |

Log files are appended to, so delete old ones before a new test run:

```bash
rm *.log
```

## Regenerating the gRPC code

Only needed if `grpc/chitchat.proto` changes. Requires `protoc`, `protoc-gen-go` and `protoc-gen-go-grpc`:

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpc/chitchat.proto
```