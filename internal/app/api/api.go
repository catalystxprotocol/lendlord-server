package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lendlord/lendlord-server/configs"
	"github.com/lendlord/lendlord-server/internal/app/api/controllers"
	"github.com/lendlord/lendlord-server/internal/app/api/router"
	"github.com/lendlord/lendlord-server/internal/app/api/server"
)

func NewApiServer(configs *configs.Server, controllers *controllers.Controllers) server.ApiServer {
	addr := fmt.Sprintf("%v:%v", "0.0.0.0", configs.Port)
	r := router.NewRouter(configs.Env)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "health")
	})
	r.GET("/chain/configs", controllers.ChainConfigController.QueryChainConfigList)

	userGroup := r.Group("/user")
	userGroup.GET("/:user_addr/collections", controllers.NftController.QueryCollectionListByUser).
		GET("/:user_addr/nfts", controllers.NftController.PageQueryNftListByUser)

	nftGroup := r.Group("/nft")
	nftGroup.GET("/detail/:nft_id", controllers.NftController.QueryNftDetail).
		GET("/activities", controllers.NftController.PageQueryNftActivityList).
		GET("/lend_info/:nft_id", controllers.NftController.QueryNftLendInfo)

	homeGroup := r.Group("/home")
	homeGroup.GET("/collection_sum", controllers.NftController.QueryCollectionSum)

	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	return server.NewServer(srv)
}
