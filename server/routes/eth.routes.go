package routes

import (
	"github.com/cryptowat-go/server/controllers"
	"github.com/cryptowat-go/server/middleware"
	"github.com/cryptowat-go/server/services"
	"github.com/gin-gonic/gin"
)

type ETHRouterController struct {
	ethController controllers.ETHController
}

func NewETHController(controller controllers.ETHController) ETHRouterController {
	return ETHRouterController{controller}
}

func (eth *ETHRouterController) ETHRoute(rg *gin.RouterGroup, userService services.UserService) {
	router := rg.Group("eth")
	router.Use(middleware.DeserializeUser(userService))
	router.GET("/current-price", eth.ethController.GetCurrentPrice)
	router.GET("/positions", eth.ethController.GetPositions)
}
