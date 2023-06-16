package app

import (
	"time"

	"github.com/oklog/oklog/pkg/group"
	"github.com/lendlord/lendlord-server/configs"
	"github.com/lendlord/lendlord-server/internal/pkg/handler"
	"github.com/lendlord/lendlord-server/internal/pkg/initialize"
)

func Server(config *configs.Config) {
	utcZone := time.FixedZone("UTC", 0)
	time.Local = utcZone
	log := initialize.InitLogger(&config.Server)

	gormDB, err := initialize.InitGormDB(config.Mysql, log)
	if err != nil {
		log.Errorf("init gormDB err: %v", err.Error())
		return
	}
	log.Debug("Init Mysql Success!")

	repo := initialize.InitRepo(gormDB, log)
	log.Debug("InitRepo Success!")

	services := initialize.InitServices(repo)
	log.Debug("InitServices Success!")

	controllers := initialize.InitControllers(services)
	log.Debug("InitControllers Success!")

	g := &group.Group{}
	handler.InitApiHandler(g, &config.Server, controllers, log)
	log.Debug("InitApiHandler Success!")
	handler.InitMetricHandler(g, &config.Server, log)
	handler.InitCancelInterrupt(g)
	log.Fatal(g.Run())
}
