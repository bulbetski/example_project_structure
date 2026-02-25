package health

import (
	"context"
	"errors"
	"net"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"

	health "github.com/builbetski/example_project_structure/gen/servergrpc/health"
)

type fakeGRPCService struct {
	err error
}

func (f *fakeGRPCService) Ping(_ context.Context) error {
	return f.err
}

func dialer(t *testing.T, srv *grpc.Server) (*grpc.ClientConn, func()) {
	listener := bufconn.Listen(1024 * 1024)
	go func() {
		_ = srv.Serve(listener)
	}()

	conn, err := grpc.DialContext(t.Context(), "bufnet", grpc.WithContextDialer(func(_ context.Context, _ string) (net.Conn, error) {
		return listener.Dial()
	}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)

	cleanup := func() {
		_ = conn.Close()
		srv.Stop()
		_ = listener.Close()
	}

	return conn, cleanup
}

func TestServer_Check_OK(t *testing.T) {
	grpcServer := grpc.NewServer()
	health.RegisterHealthServer(grpcServer, NewServer(&fakeGRPCService{}))

	conn, cleanup := dialer(t, grpcServer)
	defer cleanup()

	client := health.NewHealthClient(conn)
	resp, err := client.Check(t.Context(), &health.HealthCheckRequest{})

	require.NoError(t, err)
	require.Equal(t, health.HealthCheckResponse_SERVING, resp.Status)
}

func TestServer_Check_Error(t *testing.T) {
	grpcServer := grpc.NewServer()
	health.RegisterHealthServer(grpcServer, NewServer(&fakeGRPCService{err: errors.New("boom")}))

	conn, cleanup := dialer(t, grpcServer)
	defer cleanup()

	client := health.NewHealthClient(conn)
	resp, err := client.Check(t.Context(), &health.HealthCheckRequest{})

	require.NoError(t, err)
	require.Equal(t, health.HealthCheckResponse_NOT_SERVING, resp.Status)
}
