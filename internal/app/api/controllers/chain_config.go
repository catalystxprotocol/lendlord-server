package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lendlord/lendlord-server/internal/app/api/response"
	"github.com/lendlord/lendlord-server/internal/app/services"
)

type ChainConfigController struct {
	chainConfigService services.IChainConfigService
}

func NewChainConfigController(chainConfigService services.IChainConfigService) *ChainConfigController {
	return &ChainConfigController{
		chainConfigService: chainConfigService,
	}
}

func (ctrl *ChainConfigController) QueryChainConfigList(c *gin.Context) {
	res, err := ctrl.chainConfigService.QueryChainConfigs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailError(err))
		return
	}
	c.JSON(http.StatusOK, response.Success(res))
}
