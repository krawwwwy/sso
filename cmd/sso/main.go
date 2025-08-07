package main

import (
	"log/slog"
	"os"
	"os/signal"
	"sso/internal/app"
	"sso/internal/config"
	"sso/internal/lib/logger/slogpretty"
	"syscall"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() { // точка входа в приложение
	// TODO : init config'

	cfg := config.MustLoad()

	// TODO : init logger

	log := setupLogger(cfg.Env)
	log = log.With(
		slog.String("env", cfg.Env),
	)

	log.Debug("cfg is successfully load", slog.String("config_path", os.Getenv("CONFIG_PATH")))

	// TODO : init app

	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)

	go application.GRPCSrv.MustRun()

	// Gracefull shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCSrv.Stop()

	// TODO : lounch gRPC-server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	log = setupPrettySlog(env)
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
