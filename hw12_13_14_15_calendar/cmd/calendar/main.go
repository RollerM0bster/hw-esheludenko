package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	storage2 "github.com/RollerM0bster/hw-esheludenko/hw12_13_14_15_calendar/internal/storage"

	config2 "github.com/RollerM0bster/hw-esheludenko/hw12_13_14_15_calendar/internal/config"

	"github.com/RollerM0bster/hw-esheludenko/hw12_13_14_15_calendar/internal/app"
	"github.com/RollerM0bster/hw-esheludenko/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/RollerM0bster/hw-esheludenko/hw12_13_14_15_calendar/internal/server/http"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./configs/calendar-config.yaml", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}
	if err := run(); err != nil {
		log.Printf("%+v", err)
		os.Exit(1)
	}
}

func run() error {
	if configFile == "" {
		return fmt.Errorf("mssing configuration file")
	}
	config := config2.NewConfig()
	if err := config.Load(configFile); err != nil {
		return fmt.Errorf("error loading configuration file: %w", err)
	}
	log := logger.New(config.Logger.Level)
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	storage, err := storage2.NewStorage(ctx, config)
	if err != nil {
		return fmt.Errorf("error initializing storage: %w", err)
	}
	calendar := app.New(log, storage)
	server := internalhttp.NewServer(log, calendar)
	go shutDown(ctx, log, server)
	log.Info("calendar is running...")
	if err := server.Start(ctx, config); err != nil {
		return fmt.Errorf("error starting http server: %w", err)
	}
	return nil
}

func shutDown(ctx context.Context, log *logger.Logger, server *internalhttp.Server) {
	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	if err := server.Stop(shutdownCtx); err != nil {
		log.Error("failed to stop http server: " + err.Error())
	}
}
