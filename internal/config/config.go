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
	DB     DB
}

type DB struct {
	HostPort string
	User     string
	Password string
	DBName   string
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

	db, err := getDB(v)
	if err != nil {
		return Config{}, err
	}

	server, err := getServer(v)
	if err != nil {
		return Config{}, err
	}

	return Config{
		AppVersion:  v.GetString(appVersionKey),
		ServiceName: v.GetString(serviceNameKey),
		Server:      server,
		DB:          db,
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

func getDB(v *viper.Viper) (DB, error) {
	const (
		hostPortKey = "DB_HOST_PORT"
		userKey     = "DB_USER"
		passwordKey = "DB_PASSWORD"
		nameKey     = "DB_NAME"
	)

	var db DB

	if !v.IsSet(hostPortKey) {
		return db, fmt.Errorf("%w: %s", ErrMissingRequiredConfig, hostPortKey)
	}
	if !v.IsSet(userKey) {
		return db, fmt.Errorf("%w: %s", ErrMissingRequiredConfig, userKey)
	}
	if !v.IsSet(passwordKey) {
		return db, fmt.Errorf("%w: %s", ErrMissingRequiredConfig, passwordKey)
	}
	if !v.IsSet(nameKey) {
		return db, fmt.Errorf("%w: %s", ErrMissingRequiredConfig, nameKey)
	}

	db.HostPort = v.GetString(hostPortKey)
	db.User = v.GetString(userKey)
	db.Password = v.GetString(passwordKey)
	db.DBName = v.GetString(nameKey)

	return db, nil
}
