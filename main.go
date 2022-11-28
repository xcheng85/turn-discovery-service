package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/goava/di"
	"github.com/xcheng85/turn-discovery-service/controllers"
	"github.com/xcheng85/turn-discovery-service/middlewares"
	"github.com/xcheng85/turn-discovery-service/utils"
	"go.uber.org/zap"
)

func main() {
	di.SetTracer(&di.StdTracer{})
	// create container
	c, err := di.New(
		di.Provide(NewContext),  // provide application context
		di.Provide(NewServer),   // provide http server
		di.Provide(NewServeMux), // provide http serve mux
		di.Provide(utils.NewLogger),
		di.Provide(utils.NewConfig),
		// controllers as []Controller group
		di.Provide(controllers.NewK8sLivenessProbeController, di.As(new(controllers.Controller))),
		di.Provide(controllers.NewK8sReadinessProbeController, di.As(new(controllers.Controller))),
		di.Provide(controllers.NewTurnController, di.As(new(controllers.Controller))),
	)
	// handle container errors
	if err != nil {
		log.Fatal(err)
	}
	// invoke function
	if err := c.Invoke(StartServer); err != nil {
		log.Fatal(err)
	}
}

// StartServer starts http server.
func StartServer(ctx context.Context, server *http.Server) error {
	log.Println("start server")
	errChan := make(chan error)
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
	}()
	select {
	case <-ctx.Done():
		log.Println("stop server")
		return server.Close()
	case err := <-errChan:
		return fmt.Errorf("server error: %s", err)
	}
}

// NewContext creates new application context.
func NewContext() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		stop := make(chan os.Signal)
		signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
		<-stop
		cancel()
	}()
	return ctx
}

// NewServer creates a http server with provided mux as handler.
func NewServer(mux *http.ServeMux, logger *zap.SugaredLogger) *http.Server {
	port := utils.GetEnvVar("PORT", false, "8080")
	server := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		Handler: middlewares.MiddlewareManager(mux, logger, middlewares.LogRequest, middlewares.SecureResponse,
			middlewares.Cors),
	}
	return server
}

// NewServeMux creates a new http serve mux.
func NewServeMux(controllers []controllers.Controller, cfg *utils.AppConfig) *http.ServeMux {
	mux := &http.ServeMux{}
	for _, controller := range controllers {
		controller.RegisterRoutes(mux, cfg)
	}
	return mux
}
