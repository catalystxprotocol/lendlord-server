package router

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/lendlord/lendlord-server/internal/app/api/middleware"
	"github.com/lendlord/lendlord-server/tools"
)

func NewRouter(env string) *gin.Engine {
	if env == "prod" {
		gin.SetMode(gin.ReleaseMode)
		err := tools.CreateFolder("./logs")
		if err != nil {
			gin.DefaultWriter = os.Stdout
			log.Printf("create logs folder fail, err is %v\n", err)
		} else {
			logFilePath := fmt.Sprintf("./%v/%v", "logs", "rs_access.log")
			logFile, _ := os.OpenFile(logFilePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
			gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
		}
		gin.DisableConsoleColor()
	}

	if env != "prod" {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger())
	log.SetOutput(gin.DefaultWriter)
	middleware.SetCors(r)
	return r
}
