package relay

import (
	"context"

	api "github.com/troydai/grpcrelay/api/protos"
)

type RelayServer struct {
	api.UnimplementedRelayServer
}

func NewServer() *RelayServer {
	return &RelayServer{}
}

func (s *RelayServer) Forward(context.Context, *api.ForwardReqeust) (*api.ForwardResponse, error) {
	return nil, nil
}
