package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"prog-grpcls/internal/app"
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

	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)

	go application.GRPCSrv.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sig := <-stop

	log.Info("stopping application", slog.String("signal", sig.String()))
	application.GRPCSrv.Stop()
	log.Info("application stopped!")
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
