package main

import (
	"GoURLShortener/internal/config"
	"GoURLShortener/internal/lib/logger/sl"
	"GoURLShortener/internal/storage/sqlite"
	"fmt"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.ConfLoad()

	fmt.Printf("cfg.Env = %s\n", cfg.Env)

	log := setupLogger(cfg.Env)
	log.Info("Starting GoUrlShortener", slog.String("env", cfg.Env))
	log.Debug("Debug message for test")

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	id, err := storage.SaveURL("http://doloresl.ru", "doloresl")
	if err != nil {
		log.Error("failed to save url", sl.Err(err))
		os.Exit(1)
	}
	log.Info("saved url", slog.Int64("id", id))

	id, err = storage.SaveURL("http://doloresl.ru", "doloresl")
	if err != nil {
		log.Error("failed to save url", sl.Err(err))
		os.Exit(1)
	}

	_ = storage

	fmt.Println(cfg)
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
