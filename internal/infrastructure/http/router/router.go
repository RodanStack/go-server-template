package router

import (
	"go-server-template/internal/pkg/apires"
	"go-server-template/pkg/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Router struct {
	*gin.Engine
}

type Dependencies struct {
	fx.In

	Env *config.Env
}

func NewRouter(deps Dependencies) *Router {
	environment := deps.Env.Environment

	switch environment {
	case "production":
		gin.SetMode(gin.ReleaseMode)
	case "development":
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()

	const maxMultipartMemory = 10 << 20 // 10 MiB

	router.MaxMultipartMemory = maxMultipartMemory

	if err := router.SetTrustedProxies(nil); err != nil {
		panic(err)
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	router.GET("/health", func(c *gin.Context) {
		apires.Success(c, "OK", nil)
	})

	return &Router{
		router,
	}
}
