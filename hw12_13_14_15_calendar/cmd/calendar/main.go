package main

import (
	"context"
	"flag"
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

	if configFile == "" {
		log.Fatal("Missing configuration file")
	}
	config := config2.NewConfig()
	err := config.Load(configFile)
	if err != nil {
		log.Fatalf("Error loading configuration file: %s", err)
		os.Exit(1) //nolint:gocritic
	}
	logg := logger.New(config.Logger.Level)
	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	storage, err := storage2.NewStorage(ctx, config)
	if err != nil {
		log.Fatal(err)
	}
	calendar := app.New(logg, storage)

	server := internalhttp.NewServer(logg, calendar)

	go func() {
		<-ctx.Done()

		ctx, cancel = context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err = server.Stop(ctx); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}
	}()

	logg.Info("calendar is running...")

	if err = server.Start(ctx, config); err != nil {
		logg.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}
