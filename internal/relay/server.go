package relay

import (
	"context"

	api "github.com/troydai/grpcrelay/api/protos/relay"
	"github.com/troydai/grpcrelay/internal/settings"
)

type RelayServer struct {
	api.UnimplementedRelayServer

	config settings.Config
}

func NewServer(config settings.Config) *RelayServer {
	return &RelayServer{
		config: config,
	}
}

func (s *RelayServer) Forward(context.Context, *api.ForwardReqeust) (*api.ForwardResponse, error) {
	return nil, nil
}
