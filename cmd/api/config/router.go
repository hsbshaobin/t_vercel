package config

import (
	"base_go_web/cmd/api/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(debug bool) *gin.Engine {

	r := gin.New()
	gin.SetMode(gin.ReleaseMode)
	r.Use(func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		}
	})
	gin.SetMode(gin.ReleaseMode)

	api := r.Group("/content")

	server := api.Group("/group1")
	{
		server.GET("/test", handler.Test)

	}

	return r
}
