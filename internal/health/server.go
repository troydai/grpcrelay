package health

import (
	context "context"
	"fmt"

	"github.com/troydai/grpcrelay/api/protos/health"
)

var _ health.HealthServer = (*HealthServer)(nil)

// HeathServer implements the grpc health check service.
type HealthServer struct {
	health.UnimplementedHealthServer
}

func NewServer() *HealthServer {
	return &HealthServer{}
}

// Check implements health.HealthServer
func (*HealthServer) Check(context.Context, *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{
		Status: health.HealthCheckResponse_SERVING,
	}, nil
}

// Watch implements health.HealthServer
func (*HealthServer) Watch(req *health.HealthCheckRequest, server health.Health_WatchServer) error {
	return fmt.Errorf("not supported")
}
