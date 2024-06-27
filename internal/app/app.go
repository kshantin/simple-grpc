package app

import (
	"log/slog"
	"time"

	grpcapp "prog-grpcls/internal/app/grpc"
	"prog-grpcls/internal/services/auth"
	"prog-grpcls/internal/storage/sqlite"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}
	authService := auth.New(log, storage, storage, storage, tokenTTL)

	grpcapp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCSrv: grpcapp,
	}
}
