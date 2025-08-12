package main

import (
	"log"

	"ithub.com/liangshaoze/ops-platform/go-backend/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置和数据库...

	r := gin.Default()

	// 静态文件服务
	r.Static("/images", "./static/images")

	// 注册路由
	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		routes.RegisterAuthRoutes(auth)
	}

	// 启动服务
	log.Println("Server running on :8080")
	r.Run(":8080")
}
