package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/cryptowat-go/server/config"
	"github.com/cryptowat-go/server/controllers"
	"github.com/cryptowat-go/server/models"
	"github.com/cryptowat-go/server/routes"
	"github.com/cryptowat-go/server/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	cfg    config.Config
	server *gin.Engine
	ctx    context.Context
	db     *gorm.DB
	err    error
	//redisclient *redis.Client

	userService         services.UserService
	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	authService            services.AuthService
	AuthController         controllers.AuthController
	AuthRouteController    routes.AuthRouteController
	SessionRouteController routes.SessionRouteController

	ETHService         services.ETHServices
	ETHController      controllers.ETHController
	ETHRouteController routes.ETHRouterController
)

func init() {
	cfg, err = config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	ctx = context.TODO()

	// Connect to PostgresSql
	db = models.OpenDb(cfg)
	if err != nil {
		panic(err)
	}

	fmt.Println("Postgrest successfully connected...")

	// Collections
	userService = services.NewUserServiceImpl(db, ctx)
	authService = services.NewAuthService(db, ctx)
	ETHService = services.NewEthService(cfg, db)
	AuthController = controllers.NewAuthController(authService, userService)
	AuthRouteController = routes.NewAuthRouteController(AuthController)
	SessionRouteController = routes.NewSessionRouteController(AuthController)

	UserController = controllers.NewUserController(userService)
	UserRouteController = routes.NewRouteUserController(UserController)

	ETHController = controllers.NewETHController(ETHService)
	ETHRouteController = routes.NewETHController(ETHController)

	server = gin.Default()
	server.Use(static.Serve("/", static.LocalFile("./dist", true)))

	ETHService.InitWebSocket()
}

func main() {

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", "http://localhost:3000", "https://cryptowat-app.herokuapp.com", "*"}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))
	server.LoadHTMLGlob("dist/*.html")
	frontendRoutes := []string{
		"/",
		"/profile",
		"/login",
		"/logout",
		"/eth",
		"eth/positions",
	}

	for _, route := range frontendRoutes {
		server.GET(route, Home)
	}
	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "OK"})
	})

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router, userService)
	SessionRouteController.SessionRoute(router)
	ETHRouteController.ETHRoute(router, userService)
	log.Fatal(server.Run(":" + cfg.Port))
}

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
