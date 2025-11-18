package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"my-go-backend/internal/search/config"
	"my-go-backend/internal/search/engine"
	"my-go-backend/internal/search/httpserver"
	"my-go-backend/internal/search/service"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	client, err := engine.NewBleveClient(engine.Options{
		Directory:      cfg.IndexDir,
		AllowedIndexes: cfg.AllowedIndexes,
		SuggestFields:  cfg.SuggestFields,
	})
	if err != nil {
		log.Fatalf("failed to open search indexes: %v", err)
	}
	defer func() {
		if err := client.Close(); err != nil {
			log.Printf("search engine close warning: %v", err)
		}
	}()

	svc := service.New(client, cfg.AllowedIndexes, cfg.EnableMetrics)
	srv := httpserver.New(cfg.HTTPAddr, cfg.MetricsAddr, cfg.EnableMetrics, svc)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	log.Printf("search api starting on %s (metrics %s, indexes=%d)", cfg.HTTPAddr, cfg.MetricsAddr, len(cfg.AllowedIndexes))
	if err := srv.Run(ctx); err != nil {
		log.Fatalf("search api terminated with error: %v", err)
	}
}
