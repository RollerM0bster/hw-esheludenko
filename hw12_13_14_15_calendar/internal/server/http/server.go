package internalhttp

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/RollerM0bster/hw-esheludenko/hw12_13_14_15_calendar/internal/config"

	"github.com/RollerM0bster/hw-esheludenko/hw12_13_14_15_calendar/internal/app"

	"github.com/RollerM0bster/hw-esheludenko/hw12_13_14_15_calendar/internal/logger"
)

type Server struct {
	logger  *logger.Logger
	server  *http.Server
	app     *app.App
	wg      sync.WaitGroup
	stopped bool
	mu      sync.Mutex
}

func NewServer(logger *logger.Logger, app *app.App) *Server {
	return &Server{
		logger: logger,
		app:    app,
	}
}

func (s *Server) Start(ctx context.Context, cfg config.Config) error {
	s.logger.Info("Starting server")
	mux := http.NewServeMux()
	mux.HandleFunc("/hello-world", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Hello World"))
	})
	handler := loggingMiddleware(mux)

	s.server = &http.Server{
		Addr:              cfg.ServerConfig.Host + ":" + cfg.ServerConfig.Port,
		Handler:           handler,
		ReadHeaderTimeout: 10 * time.Second,
	}
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		if err := s.server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			s.logger.Error("Server error: " + err.Error())
		}
	}()

	<-ctx.Done()
	return s.Stop(ctx)
}

func (s *Server) Stop(_ context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.stopped {
		return nil
	}
	s.stopped = true
	s.logger.Info("Stopping server")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.server.Shutdown(shutdownCtx); err != nil {
		s.logger.Error("Server error: " + err.Error())
		return err
	}
	s.wg.Wait()
	s.logger.Info("Server stopped")
	return nil
}
