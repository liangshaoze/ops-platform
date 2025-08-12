package routes

import (
	"ithub.com/liangshaoze/ops-platform/go-backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.RouterGroup) {
	// 验证码
	r.GET("/captcha", handlers.GenerateCaptcha)

	// 登录
	r.POST("/login", handlers.Login)

	// 第三方登录
	r.GET("/wecom", handlers.WeComAuth)
	r.GET("/dingtalk", handlers.DingTalkAuth)
}
