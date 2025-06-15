package main

import (
	"context"

	gokit "github.com/cripplemymind9/go-utils/go-kit"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/cripplemymind9/inventory-service/internal/app"
	"github.com/cripplemymind9/inventory-service/internal/config"
)

func main() {
	cfg, err := config.Get(viper.New())
	if err != nil {
		log.Fatal().Err(err).Msg("get config")
	}

	runner := gokit.NewRunner()

	app, err := app.New(context.Background(), cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("get new app")
	}

	if err := runner.Run(app); err != nil { //nolint:govet // Допускаем shadowing err для краткости
		log.Fatal().Err(err).Msg("run app")
	}
}
