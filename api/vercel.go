// api/vercel.go
package main

import (
	"base_go_web/cmd/api/config"
	"net/http"
	"strings"
	// 替换为你的项目模块名！！！
	// 查看项目根目录 go.mod 文件第一行：module base_go_web
)

// Vercel 强制要求的入口函数：Handler
func Handler(w http.ResponseWriter, r *http.Request) {
	// 获取你原有 Gin 引擎（二选一）
	ginEngine := config.GetVercelGinEngine() // 方案A
	// ginEngine := api.GetGinEngine() // 方案B
	if strings.HasPrefix(r.URL.Path, "/api/vercel") {
		newPath := strings.TrimPrefix(r.URL.Path, "/api/vercel")
		if newPath == "" {
			newPath = "/"
		}
		r.URL.Path = newPath
	}

	// 将 Gin 引擎适配为 http.Handler 处理请求（核心！）
	ginEngine.ServeHTTP(w, r)
}
