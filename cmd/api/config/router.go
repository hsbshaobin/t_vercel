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
		// 允许所有前端域名访问（小游戏部署后域名）
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

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

// 新增：导出供 Vercel 使用的 Gin 引擎（复用原有 InitRouter）
func GetVercelGinEngine() *gin.Engine {
	// 复用你原有路由初始化逻辑，debug 传 true/false 都可以
	return InitRouter(true)
}
