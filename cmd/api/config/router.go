package config

import (
	"adCheck/cmd/api/handler"
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

	api := r.Group("/ac")

	server := api.Group("/check")
	{
		//server.POST("/generateCode", handler.GenerateCode)
		server.POST("/checkCode", handler.CheckCode)

	}

	return r
}
