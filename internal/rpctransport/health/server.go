package health

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/builbetski/example_project_structure/gen/servergrpc/health"
)

type HealthService interface {
	Ping(ctx context.Context) error
}

type Server struct {
	health.UnimplementedHealthServer
	svc HealthService
}

func NewServer(svc HealthService) *Server {
	return &Server{svc: svc}
}

func (s *Server) Check(ctx context.Context, _ *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	if err := s.svc.Ping(ctx); err != nil {
		return &health.HealthCheckResponse{Status: health.HealthCheckResponse_NOT_SERVING}, nil
	}

	return &health.HealthCheckResponse{Status: health.HealthCheckResponse_SERVING}, nil
}

func (s *Server) Watch(_ *health.HealthCheckRequest, _ health.Health_WatchServer) error {
	return status.Error(codes.Unimplemented, "watch is not implemented")
}
