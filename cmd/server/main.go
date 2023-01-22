package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	healthapi "github.com/troydai/grpcrelay/api/protos/health"
	relayapi "github.com/troydai/grpcrelay/api/protos/relay"
	"github.com/troydai/grpcrelay/internal/health"
	"github.com/troydai/grpcrelay/internal/relay"
	"github.com/troydai/grpcrelay/internal/settings"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	configLoader, err := settings.NewFileConfigLoader(settings.WithAllowFileMissing())
	if err != nil {
		log.Fatal(err)
	}

	c, err := configLoader.Load(ctx)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	rs := relay.NewServer(c)
	hs := health.NewServer()

	relayapi.RegisterRelayServer(server, rs)
	healthapi.RegisterHealthServer(server, hs)
	reflection.Register(server)

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(fmt.Errorf("fail to start TCP listener: %w", err))
	}

	chServerStopped := make(chan struct{})
	chSystemSignal := make(chan os.Signal, 1)

	signal.Notify(chSystemSignal)

	go func() {
		select {
		case <-chServerStopped:
		case <-chSystemSignal:
			server.GracefulStop()
		}
	}()

	go func() {
		defer close(chServerStopped)
		server.Serve(lis)
		fmt.Println("Gracefully stopped")
	}()

	<-chServerStopped
}
