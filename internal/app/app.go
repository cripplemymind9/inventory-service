package app

import (
	"context"
	"errors"
	"fmt"
	"time"

	gokit "github.com/cripplemymind9/go-utils/go-kit"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"github.com/cripplemymind9/inventory-service/internal/config"
	"github.com/cripplemymind9/inventory-service/internal/server"
)

type App struct {
	gokit.App

	ctx    context.Context
	cancel context.CancelFunc

	cfg config.Config

	server *server.Server
}

func (a *App) Run() error {
	<-a.ctx.Done()
	log.Info().Msg("bye")

	if errors.Is(a.ctx.Err(), context.Canceled) {
		return fmt.Errorf("context cancelled run app err: %w", a.ctx.Err())
	}

	return nil
}

func (a *App) Shutdown(dur time.Duration) error {
	time.Sleep(dur)
	a.cancel()

	return nil
}

func (a *App) RegisterGRPCServices(server grpc.ServiceRegistrar) {
	a.server.RegisterServices(server)
}

func (a *App) RegisterHandlersFromEndpoint(
	ctx context.Context,
	mux *runtime.ServeMux,
	endpoint string,
	opts []grpc.DialOption,
) error {
	return a.server.RegisterHandlersFromEndPoint(ctx, mux, endpoint, opts)
}

func New(ctx context.Context, cfg config.Config) (*App, error) {
	serverDependencies, err := getGRPCServerDependencies(cfg)
	if err != nil {
		return nil, err
	}

	server := server.New(cfg, serverDependencies)

	ctx, cancel := context.WithCancel(ctx)

	return &App{
		ctx:    ctx,
		cancel: cancel,
		server: server,
		cfg:    cfg,
	}, nil
}

// nolint:unparam // Функция возвращает error для соответствия интерфейсу, но в текущей реализации ошибка всегда nil
func getGRPCServerDependencies(_ config.Config) (*server.Dependencies, error) {
	return server.NewDependencies(), nil
}
