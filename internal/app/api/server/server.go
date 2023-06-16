package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	logger "github.com/sirupsen/logrus"
)

type ApiServer interface {
	Start() error
}

type apiServer struct {
	srv *http.Server
}

func NewServer(server *http.Server) ApiServer {
	return &apiServer{srv: server}
}

func (api *apiServer) Start() error {
	go func() {
		if err := api.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("listen server err: %v", err.Error())
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Info("Shutdown Server...")

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := api.srv.Shutdown(ctx); err != nil {
		logger.Errorf("Server Shutdown: %v", err.Error())
		return err
	}
	return nil
}
