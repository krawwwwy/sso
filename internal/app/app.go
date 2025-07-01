package app

import (
	"log/slog"
	"os"
	"sso/internal/app/grpcapp"
	"sso/internal/lib/logger/sl"
	"sso/internal/service/auth"
	"sso/internal/storage/sqlite"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	storage, err := sqlite.New(storagePath)
	if err != nil {
		log.Error("storage initialization have failes", sl.Err(err))
		os.Exit(1)
	}

	authService := auth.New(log, storage, storage, storage, tokenTTL)
	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
