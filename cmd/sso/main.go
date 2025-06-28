package main

import (
	"log/slog"
	"os"
	"sso/internal/config"
	"sso/internal/lib/logger/slogpretty"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() { // точка входа в приложение
	// TODO : init config'

	cfg := config.MustLoad()
	_ = cfg

	log := setupLogger(cfg.Env)

	log.Debug("cfd is successfully load", slog.String("CONFIG_PATH", os.Getenv("CONFIG_PATH")))
	// TODO : init logger

	// TODO : init app

	// TODO : lounch gRPC-server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog(env)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	}

	return log

}

func setupPrettySlog(env string) *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}
	switch env {
	case envLocal:
		opts.SlogOpts.Level = slog.LevelDebug
	case envDev:
		opts.SlogOpts.Level = slog.LevelDebug
	case envProd:
		opts.SlogOpts.Level = slog.LevelInfo
	}
	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
