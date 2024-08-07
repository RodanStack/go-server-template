package router

import (
	"go-server-template/internal/pkg/apires"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

type Dependencies struct {
}

func NewRouter(_ Dependencies) *Router {
	router := gin.Default()

	const maxMultipartMemory = 10 << 20 // 10 MiB

	router.MaxMultipartMemory = maxMultipartMemory

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
