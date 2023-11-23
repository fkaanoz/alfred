package main

import (
	"context"
	"fmt"
	"github.com/fkaanoz/alfred/business"
	"github.com/fkaanoz/alfred/business/handlers"
	"github.com/fkaanoz/alfred/business/mids"
	"github.com/fkaanoz/alfred/foundation/server"
	"go.uber.org/zap"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var Build string
var Service string = "Alfred"

func main() {
	logger, err := business.Initlogger(Service, Build)
	if err != nil {
		log.Println("init logger err : ", err)
		os.Exit(1)
	}

	if err := run(logger); err != nil {
		fmt.Println("err", err)
		os.Exit(1)
	}
}

func run(logger *zap.SugaredLogger) error {
	port := "9090"

	mux := server.NewServer(server.Config{
		Logger:      logger,
		Timeout:     10,
		Middlewares: []server.Middleware{mids.Logger(logger)},
	})
	mux = v1(mux, logger)

	apiHost := http.Server{
		Addr:    net.JoinHostPort("localhost", port),
		Handler: mux,
	}

	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		logger.Infow("SERVER", "status", "starting", "port", port)
		if err := apiHost.ListenAndServe(); err != nil {
			logger.Errorw("SERVER", "ERROR", err)
		}
	}()

	select {
	case <-shutdownCh:
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		logger.Infow("SERVER", "status", "gracefully shutdown started")

		if err := apiHost.Shutdown(ctx); err != nil {
			logger.Errorw("SERVER", "status", "gracefully shutdown is not possible", "ERROR", err)
			apiHost.Close()
		}
	}

	return nil
}

func v1(mux *server.Server, logger *zap.SugaredLogger) *server.Server {

	mux.GET("/api/get-test", handlers.GetTestHandler)
	mux.POST("/api/post-test", handlers.PostTestHandler)
	mux.DELETE("/api/delete-test", handlers.DeleteTestHandler)
	mux.PUT("/api/put-test", handlers.PutTestHandler)

	return mux
}
