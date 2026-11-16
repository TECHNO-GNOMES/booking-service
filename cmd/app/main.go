package main

import (
	"context"
	"log"
	"log/slog"

	"github.com/nurgal1ev/booking-service/internal/app"
	"github.com/nurgal1ev/booking-service/internal/config"
	"github.com/nurgal1ev/booking-service/internal/infrastructure/postgres"
	"github.com/nurgal1ev/booking-service/internal/infrastructure/seaweedfs"
	"github.com/nurgal1ev/booking-service/internal/transport/httpv1"
)

func main() {
	cfg := config.Load()
	slog.Info("loading config")
	ctx := context.Background()

	db, err := postgres.NewDb(cfg)
	if err != nil {
		log.Fatal(err)
	}

	s3Client, err := seaweedfs.NewConnectAWS(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}

	handlers := app.InitHandlers(db.DB, cfg, s3Client)

	httpv1.StartServer(handlers, cfg)
}
