package main

import (
	"log/slog"
	"os"

	"prog-grpcls/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// TODO: инит объект конфига
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)
	log.Info("starting app...",
		slog.String("env", cfg.Env),
		slog.Any("env", cfg),
		slog.Int("env", cfg.GRPC.Port),
	)
	// TODO: инит логгера

	// TODO: инит приложения

	// TODO: запустить gRPC-сервер приложения
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
