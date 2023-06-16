package middleware

import (
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lendlord/lendlord-server/internal/app/constant"
)

func SetCors(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowMethods: []string{"GET", "POST", "PUT", "HEAD", "DELETE", "PATCH"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", constant.HeaderSignature,
			constant.HeaderAuthorization, constant.HeaderXForwardedFor, "X-Real-Ip",
			"X-Appengine-Remote-Addr", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{constant.HeaderPagination, constant.HeaderContentDisposition},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))
	r.Use(func(c *gin.Context) {
		if strings.Contains(c.FullPath(), "debug/pprof") {
			c.Next()
			return
		}
	})
}
