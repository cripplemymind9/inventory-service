package config

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/viper"
)

var ErrMissingRequiredConfig = errors.New("missing required config")

type Config struct {
	AppVersion  string
	ServiceName string

	Server Server
}

type Server struct {
	GRPCPort        int
	HTTPPort        int
	ShutDownTimeout time.Duration
}

func Get(v *viper.Viper) (Config, error) {
	v.AutomaticEnv()

	const (
		appVersionKey  = "VERSION"
		serviceNameKey = "SERVICE_NAME"
	)

	if !v.IsSet(appVersionKey) {
		return Config{}, fmt.Errorf("%w: %s", ErrMissingRequiredConfig, appVersionKey)
	}

	if !v.IsSet(serviceNameKey) {
		return Config{}, fmt.Errorf("%w: %s", ErrMissingRequiredConfig, serviceNameKey)
	}

	server, err := getServer(v)
	if err != nil {
		return Config{}, err
	}

	return Config{
		AppVersion:  v.GetString(appVersionKey),
		ServiceName: v.GetString(serviceNameKey),
		Server:      server,
	}, nil
}

func getServer(v *viper.Viper) (Server, error) {
	const (
		grpcPortKey        = "GRPC_PORT"
		httpPortKey        = "HTTP_PORT"
		shutdownTimeoutKey = "SHUTDOWN_TIMEOUT"
	)

	var server Server

	if !v.IsSet(grpcPortKey) {
		return server, fmt.Errorf("%w: %s", ErrMissingRequiredConfig, grpcPortKey)
	}

	if !v.IsSet(httpPortKey) {
		return server, fmt.Errorf("%w: %s", ErrMissingRequiredConfig, httpPortKey)
	}

	if !v.IsSet(shutdownTimeoutKey) {
		return server, fmt.Errorf("%w: %s", ErrMissingRequiredConfig, shutdownTimeoutKey)
	}

	server.GRPCPort = v.GetInt(grpcPortKey)
	server.HTTPPort = v.GetInt(httpPortKey)
	server.ShutDownTimeout = time.Duration(v.GetInt(shutdownTimeoutKey)) * time.Second

	return server, nil
}
