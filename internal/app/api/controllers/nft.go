package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lendlord/lendlord-server/internal/app/api/response"
	"github.com/lendlord/lendlord-server/internal/app/constant"
	"github.com/lendlord/lendlord-server/internal/app/models/form"
	"github.com/lendlord/lendlord-server/internal/app/services"
)

type NftController struct {
	nftService services.INftService
}

func NewNftController(nftService services.INftService) *NftController {
	return &NftController{
		nftService: nftService,
	}
}

func (ctrl *NftController) QueryCollectionListByUser(c *gin.Context) {
	param := form.QueryCollectionListByUserForm{UserAddr: c.Param(constant.UserAddr)}
	err := c.ShouldBind(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.FailFormBindRequest(err))
		return
	}
	res, err1 := ctrl.nftService.QueryNftCollectionListByUser(&param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailError(err1))
		return
	}
	c.JSON(http.StatusOK, response.Success(res))
}

func (ctrl *NftController) PageQueryNftListByUser(c *gin.Context) {
	param := form.PageQueryNftListByUserForm{UserAddr: c.Param(constant.UserAddr)}
	err := c.ShouldBind(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.FailFormBindRequest(err))
		return
	}

	res, err1 := ctrl.nftService.PageQueryNftListByUser(&param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailError(err1))
		return
	}
	c.JSON(http.StatusOK, response.Success(res))
}

func (ctrl *NftController) QueryNftDetail(c *gin.Context) {
	nftId, err := strconv.Atoi(c.Param(constant.NftId))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.FailFormBindRequest(err))
		return
	}

	param := form.QueryNftDetailForm{NftId: nftId}
	err1 := c.ShouldBind(&param)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, response.FailFormBindRequest(err1))
		return
	}

	res, err2 := ctrl.nftService.QueryNftDetail(&param)
	if err2 != nil {
		c.JSON(err2.StatusCode(), response.FailError(err2))
		return
	}
	c.JSON(http.StatusOK, response.Success(res))
}

func (ctrl *NftController) PageQueryNftActivityList(c *gin.Context) {
	// Query Variables
	var param form.PageQueryNftActivityListForm
	err := c.ShouldBind(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.FailFormBindRequest(err))
		return
	}

	res, err1 := ctrl.nftService.PageQueryNftActivityList(&param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailError(err1))
		return
	}
	c.JSON(http.StatusOK, response.Success(res))
}

func (ctrl *NftController) QueryNftLendInfo(c *gin.Context) {
	nftId, err := strconv.Atoi(c.Param(constant.NftId))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.FailFormBindRequest(err))
		return
	}

	param := form.QueryNftLendInfoForm{NftId: nftId}
	err1 := c.ShouldBind(&param)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, response.FailFormBindRequest(err1))
		return
	}

	res, err2 := ctrl.nftService.QueryNftLendInfo(&param)
	if err2 != nil {
		c.JSON(err2.StatusCode(), response.FailError(err2))
		return
	}
	c.JSON(http.StatusOK, response.Success(res))
}

func (ctrl *NftController) QueryCollectionSum(c *gin.Context) {
	res, err := ctrl.nftService.QueryNftCollectionSum()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailError(err))
		return
	}
	c.JSON(http.StatusOK, response.Success(res))
}
