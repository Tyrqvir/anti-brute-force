package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Tyrqvir/anti-brute-force/internal/config"
	appLogger "github.com/Tyrqvir/anti-brute-force/internal/logger"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/etc/app/config.toml", "Path to configuration file")
}

func main() {
	flag.Parse()

	cfg, err := config.NewConfig(configFile)
	if err != nil {
		log.Fatalln(err)
	}

	logger := appLogger.New(cfg.Logger.Level)

	server, err := setupWire(cfg, logger)
	if err != nil {
		log.Fatalln(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer stop()

	go func() {
		if err := server.Start(ctx); err != nil {
			logger.Errorf("failed to start grpс server: %w", err)
			stop()
			os.Exit(1)
		}
	}()

	go func() {
		<-ctx.Done()

		server.Stop()
	}()

	<-ctx.Done()
}
