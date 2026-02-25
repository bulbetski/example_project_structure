package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/builbetski/example_project_structure/gen/servergrpc/health"
	"github.com/builbetski/example_project_structure/internal/cli/deps"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("app: %v", err)
	}
}

func run() error {
	container := deps.NewContainer()

	cfg := container.GetConfig()
	pool := container.GetDatabase()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	e := echo.New()
	httpHandler := container.GetHealthHTTPHandler()
	httpHandler.RegisterRoutes(e)

	httpServer := &http.Server{
		Addr:    cfg.HTTPAddr,
		Handler: e,
	}

	grpcServer := grpc.NewServer()
	rpcHealth := container.GetHealthGRPCServer()
	health.RegisterHealthServer(grpcServer, rpcHealth)
	reflection.Register(grpcServer)

	httpListener, err := net.Listen("tcp", cfg.HTTPAddr)
	if err != nil {
		return err
	}

	grpcListener, err := net.Listen("tcp", cfg.GRPCAddr)
	if err != nil {
		return err
	}

	group, groupCtx := errgroup.WithContext(ctx)
	group.Go(func() error {
		if err := httpServer.Serve(httpListener); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	})
	group.Go(func() error {
		if err := grpcServer.Serve(grpcListener); err != nil {
			return err
		}
		return nil
	})
	group.Go(func() error {
		<-groupCtx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
		defer cancel()
		_ = httpServer.Shutdown(shutdownCtx)
		grpcServer.GracefulStop()
		pool.Close()
		return nil
	})

	log.Printf("http listening on %s", cfg.HTTPAddr)
	log.Printf("grpc listening on %s", cfg.GRPCAddr)
	return group.Wait()
}
