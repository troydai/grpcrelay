package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	api "github.com/troydai/grpcrelay/api/protos"
	"github.com/troydai/grpcrelay/internal/relay"
)

func main() {
	server := grpc.NewServer()
	rs := relay.NewServer()

	api.RegisterRelayServer(server, rs)
	reflection.Register(server)

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(fmt.Errorf("fail to start TCP listener: %w", err))
	}

	server.Serve(lis)
}
