package app

import (
	"log/slog"
	"time"

	grpcapp "prog-grpcls/internal/app/grpc"
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
	grpcapp := grpcapp.New(log, grpcPort)
	return &App{
		GRPCSrv: grpcapp,
	}
}
