package handler

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/oklog/oklog/pkg/group"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"github.com/lendlord/lendlord-server/configs"
	"github.com/lendlord/lendlord-server/internal/app/api"
	"github.com/lendlord/lendlord-server/internal/app/api/controllers"
)

func InitApiHandler(g *group.Group, cfg *configs.Server, controllers *controllers.Controllers, logger *log.Logger) {
	addr := fmt.Sprintf("0.0.0.0:%s", cfg.Port)

	g.Add(func() error {
		logger.WithFields(log.Fields{
			"features":    "server",
			"transport":   "http",
			"server_addr": addr,
		}).Info("start")
		server := api.NewApiServer(cfg, controllers)
		return server.Start()
	}, func(error) {
		logger.WithFields(log.Fields{"httpListener.CLose": addr})
	})
}

func InitCancelInterrupt(g *group.Group) {
	cancelInterrupt := make(chan struct{})
	g.Add(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-c:
			return fmt.Errorf("received signal %s", sig)
		case <-cancelInterrupt:
			return nil
		}
	}, func(error) {
		close(cancelInterrupt)
	})
}

func InitMetricHandler(g *group.Group, cfg *configs.Server, logger *log.Logger) {
	metricMux := http.NewServeMux()
	metricMux.Handle("/metrics", promhttp.Handler())
	addr := fmt.Sprintf("0.0.0.0:%s", cfg.PrometheusPort)
	httpListener, err := net.Listen("tcp", addr)
	if err != nil {
		logger.WithFields(log.Fields{
			"features":    "metric",
			"transport":   "http",
			"during":      "listen",
			"metric_addr": addr,
		}).Error(err)
	}

	g.Add(func() error {
		logger.WithFields(log.Fields{
			"features":    "metric",
			"transport":   "http",
			"metric_addr": addr,
		}).Info("start")
		return http.Serve(httpListener, metricMux)
	}, func(err error) {
		logger.WithFields(log.Fields{
			"features":           "metric",
			"httpListener.Close": httpListener.Close(),
		}).Error(err)
	})
}
