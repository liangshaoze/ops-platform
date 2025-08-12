package main

import (
	"log"

	"github.com/liangshaoze/ops-platform/go-backend/internal/config"
	"github.com/liangshaoze/ops-platform/go-backend/internal/handlers"
	"github.com/liangshaoze/ops-platform/go-backend/internal/middleware"
	"github.com/liangshaoze/ops-platform/go-backend/internal/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	config.LoadConfig()

	// 初始化数据库
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 设置Gin模式
	if config.AppConfig.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 创建Gin实例
	r := gin.Default()

	// 注册中间件
	r.Use(middleware.CORSMiddleware())

	// 注册路由
	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", handlers.Login)
			auth.POST("/logout", handlers.Logout)
		}

		// 需要认证的路由
		api.Use(middleware.AuthMiddleware())
		{
			api.GET("/userinfo", handlers.GetUserInfo)
		}
	}

	// 启动服务器
	log.Printf("Server running on %s", config.AppConfig.ServerAddr)
	if err := r.Run(config.AppConfig.ServerAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
