package main

import (
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/melkishengue/gotemplate/cmd/frontend/server"
	"github.com/melkishengue/gotemplate/cmd/frontend/server/handlers"
	"github.com/melkishengue/gotemplate/internal/logger"
	"github.com/melkishengue/gotemplate/pkg/utils"
)

func main() {
	// ignore error because we dont load any .env file in prod
	_ = godotenv.Load()

	logger.SetUp(utils.GetEnvOrDie("ENVIRONMENT"))

	server, err := server.NewServer(":3010", handlers.NewStatic())
	if err != nil {
		slog.Error("failed to create server", slog.Any("err", err))
		os.Exit(1)
	}

	stopped := make(chan struct{})
	go func() {
		defer close(stopped)

		slog.Info("starting server", slog.String("address", ":3010"))

		if err := server.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("server failed", slog.Any("err", err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case <-quit:
		slog.Info("Shutting down gracefully...")
	case <-stopped:
		slog.Error("Server stopped unexpectedly, shutting down...")
	}

	if err := server.Stop(10 * time.Second); err != nil {
		slog.Error("Server failed to shutdown gracefully", slog.Any("err", err))
		os.Exit(1)
	}
}
