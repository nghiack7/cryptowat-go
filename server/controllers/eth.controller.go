package controllers

import (
	"github.com/cryptowat-go/server/models"
	"github.com/cryptowat-go/server/services"
	"github.com/gin-gonic/gin"
)

type ETHController struct {
	ETHService services.ETHServices
}

func NewETHController(ethService services.ETHServices) ETHController {
	return ETHController{ethService}
}

func (eth *ETHController) GetCurrentPrice(c *gin.Context) {
	var ethResponse models.Price
	ethResponse.Value = eth.ETHService.GetCurrentPrice()
	ethResponse.Photo = "https://cryptologos.cc/logos/ethereum-eth-logo.png?v=024"
	ethResponse.Currency = "ETH/USD"
	c.JSON(200, ethResponse)
}

func (eth *ETHController) GetPositions(c *gin.Context) {
	resp, err := eth.ETHService.GetAllPositions()
	if err != nil {
		c.AbortWithStatusJSON(500, err)
	}
	c.JSON(200, resp)
}
