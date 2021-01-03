package web

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"{{cookiecutter.gomodule_uri}}/pkg/configuration"
	"{{cookiecutter.gomodule_uri}}/pkg/domain"
	"{{cookiecutter.gomodule_uri}}/pkg/infrastructure/storage/memory"
	"{{cookiecutter.gomodule_uri}}/pkg/service"
	webdemo "{{cookiecutter.gomodule_uri}}/pkg/web/demo"

	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/rs/zerolog"
)

// NewDemoRepository instantiates a storage repository according to the configuration.
func NewDemoRepository(ctx context.Context, cfg *configuration.AppConfiguration, logger *zerolog.Logger) (domain.DemoRepository, error) {
	switch cfg.Repository.Adapter {
	case "memory":
		return memory.NewDemoRepository(ctx, cfg.Repository.Options, logger)
	default:
		return nil, fmt.Errorf("unknown storage adapter: [%s]", cfg.Repository.Adapter)
	}
}

// NewDemoService fires up a demo service
func NewDemoService(r domain.DemoRepository, l *zerolog.Logger) (*service.DemoService, error) {
	return &service.DemoService{
		Repository: r,
		Logger:     l,
	}, nil
}

// NewRouter creates a mux with mounted routes and instantiates respective dependencies.
func NewRouter(ctx context.Context, cfg *configuration.AppConfiguration, logger *zerolog.Logger) *chi.Mux {
	demoRepository, err := NewDemoRepository(ctx, cfg, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Could not instantiate the demo repository")
	}

	demoService, err := NewDemoService(demoRepository, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("Could not instantiate the demo service")
	}

	r := chi.NewRouter()

	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(chimiddleware.Heartbeat("/status"))

	r.Mount("/api", webdemo.Handler{}.Routes(logger, demoService))

	return r
}

// LaunchServer starts a web server and propagates shutdown context.
func LaunchServer(cfg *configuration.AppConfiguration, logger *zerolog.Logger) error {
	var err error

	c := make(chan os.Signal, 1)
	signal.Notify(
		c,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		s := <-c
		logger.Debug().Str("syscall", s.String()).Msg("Intercepted syscall")
		cancel()
	}()

	router := NewRouter(ctx, cfg, logger)
	srv := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%d", cfg.Port),
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal().Err(err).Msg("Could not launch the web server")
		}
	}()
	logger.Printf("Starting server on port: [%d]", cfg.Port)

	<-ctx.Done()

	logger.Printf("Cleaning up the server")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if err = srv.Shutdown(ctxShutDown); err != nil {
		logger.Fatal().Err(err).Msg("Error on server shutdown")
	}

	cancel()

	logger.Printf("Server exited successfully")

	if err == http.ErrServerClosed {
		err = nil
	}
	return err
}
