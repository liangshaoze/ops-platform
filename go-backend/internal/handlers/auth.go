/*
 * @Author: 12866449444136360 liangsz@aliyun.com
 * @Date: 2025-08-12 14:01:26
 * @LastEditors: liangsz@aliyun.com liangsz@aliyun.com
 * @LastEditTime: 2025-08-12 15:23:57
 * @FilePath: \自学项目\go-backend\internal\handlers\auth.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package handlers

import (
	"net/http"
	"time"

	"github.com/liangshaoze/ops-platform/go-backend/internal/config"
	"github.com/liangshaoze/ops-platform/go-backend/internal/models"
	"github.com/liangshaoze/ops-platform/go-backend/internal/pkg/database"
	"github.com/liangshaoze/ops-platform/go-backend/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "无效的请求参数")
		return
	}

	var user models.User
	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "用户名或密码错误")
		return
	}

	if !user.CheckPassword(req.Password) {
		utils.ErrorResponse(c, http.StatusUnauthorized, "用户名或密码错误")
		return
	}

	// 生成JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"is_admin": user.IsAdmin,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.AppConfig.JWTSecret))
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "生成Token失败")
		return
	}

	utils.SuccessResponse(c, gin.H{
		"token": tokenString,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"is_admin": user.IsAdmin,
		},
	})
}

func Logout(c *gin.Context) {
	utils.SuccessResponse(c, gin.H{
		"message": "登出成功",
	})
}

func GetUserInfo(c *gin.Context) {
	claims, _ := c.Get("claims")
	utils.SuccessResponse(c, claims)
}
