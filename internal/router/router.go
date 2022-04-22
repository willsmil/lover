package router

import (
	"github.com/gin-gonic/gin"
	"github.com/willsmil/lover/internal/api"
	"github.com/willsmil/lover/internal/middleware"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(middleware.Trace())
	router.GET("/ping", api.Health)

	return router
}
