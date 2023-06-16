package initialize

import (
	"fmt"
	"net/url"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/lendlord/lendlord-server/configs"
	"github.com/lendlord/lendlord-server/internal/app/api/controllers"
	"github.com/lendlord/lendlord-server/internal/app/repository"
	"github.com/lendlord/lendlord-server/internal/app/services"
	"github.com/lendlord/lendlord-server/internal/pkg/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitLogger(cfg *configs.Server) *log.Logger {
	Logger := log.New()
	Logger.SetReportCaller(true)

	if cfg.Env == "prod" {
		Logger.SetFormatter(&log.JSONFormatter{})
	} else {
		Logger.SetFormatter(&log.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
		})
	}
	switch cfg.LogLevel {
	case "debug":
		Logger.SetLevel(log.DebugLevel)
	case "error":
		Logger.SetLevel(log.ErrorLevel)
	case "warn":
		Logger.SetLevel(log.WarnLevel)
	default:
		Logger.SetLevel(log.InfoLevel)
	}
	return Logger
}

func InitGormDB(cfg configs.Mysql, logger *log.Logger) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&time_zone=%s&max_execution_time=%d",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, url.QueryEscape("'UTC'"), cfg.MaxExecutionTime)

	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 logs.NewGormLogger(logger),
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   cfg.TablePrefix,   // 表前缀
			SingularTable: cfg.SingularTable, // 复数形式
		},
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Hour * time.Duration(cfg.MaxLifetime))

	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	return gormDB, nil
}

func InitRepo(gormDB *gorm.DB, log *log.Logger) *repository.Repos {
	return &repository.Repos{
		ChainConfigRepo:   repository.NewChainConfigRepo(gormDB),
		NftRepo:           repository.NewNftRepo(gormDB),
		NftActivityRepo:   repository.NewNftActivityRepo(gormDB),
		NftCollectionRepo: repository.NewNftCollectionRepo(gormDB),
		NftOrderRepo:      repository.NewNftOrdersRepo(gormDB),
		Log:               log,
	}
}

func InitServices(repo *repository.Repos) *services.Service {
	return &services.Service{
		ChainConfigService: services.NewChainConfigService(repo.ChainConfigRepo, repo.Log),
		NftService:         services.NewNftService(repo.NftRepo, repo.NftActivityRepo, repo.NftCollectionRepo, repo.NftOrderRepo, repo.Log),
	}
}

func InitControllers(services *services.Service) *controllers.Controllers {
	return &controllers.Controllers{
		ChainConfigController: controllers.NewChainConfigController(services.ChainConfigService),
		NftController:         controllers.NewNftController(services.NftService),
	}
}
