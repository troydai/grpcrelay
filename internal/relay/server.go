package relay

import (
	"context"
	"fmt"

	beaconapi "github.com/troydai/grpcrelay/api/protos/beacon"
	relayapi "github.com/troydai/grpcrelay/api/protos/relay"
	"github.com/troydai/grpcrelay/internal/settings"
	"google.golang.org/grpc"
)

type RelayServer struct {
	relayapi.UnimplementedRelayServer

	config settings.Config
}

func NewServer(config settings.Config) *RelayServer {
	return &RelayServer{
		config: config,
	}
}

func (s *RelayServer) Forward(ctx context.Context, req *relayapi.ForwardReqeust) (*relayapi.ForwardResponse, error) {
	for _, rec := range s.config.Recievers {
		switch rec.ServiceType {
		case "grpcrelay.Relay":
			forwawrdToRelay(ctx, rec.Address)
		case "grpcbeacon.Beacon":
			forwardToBeacon(ctx, rec.Address)
		default:
			// skip over
		}
	}

	return &relayapi.ForwardResponse{}, nil
}

func forwardToBeacon(ctx context.Context, addr string) error {
	conn, err := grpc.DialContext(ctx, addr)
	if err != nil {
		return fmt.Errorf("fail to create gRPC connection for service at %s: %w", addr, err)
	}
	defer conn.Close()

	client := beaconapi.NewBeaconClient(conn)

	_, err = client.Signal(ctx, &beaconapi.SignalReqeust{})
	if err != nil {
		return fmt.Errorf("fail to forward request to beacon service at %s: %w", addr, err)
	}

	return nil
}

func forwawrdToRelay(ctx context.Context, addr string) error {
	conn, err := grpc.DialContext(ctx, addr)
	if err != nil {
		return fmt.Errorf("fail to create gRPC connection for service at %s: %w", addr, err)
	}
	defer conn.Close()

	client := relayapi.NewRelayClient(conn)

	_, err = client.Forward(ctx, &relayapi.ForwardReqeust{})
	if err != nil {
		return fmt.Errorf("fail to forward request to relay service at %s: %w", addr, err)
	}

	return nil
}
