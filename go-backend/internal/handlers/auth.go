package handlers

import (
	"ops-platform/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

// 验证码存储
var captchaStore = base64Captcha.DefaultMemStore

// GenerateCaptcha 生成验证码
func GenerateCaptcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)
	captcha := base64Captcha.NewCaptcha(driver, captchaStore)

	id, b64s, err := captcha.Generate()
	if err != nil {
		utils.InternalServerError(c, "生成验证码失败")
		return
	}

	utils.Success(c, gin.H{
		"captcha_id": id,
		"image":      b64s,
	})
}

// VerifyCaptcha 验证验证码
func VerifyCaptcha(captchaId, captcha string) bool {
	return captchaStore.Verify(captchaId, captcha, true)
}

// Login 登录处理
func Login(c *gin.Context) {
	var req struct {
		Username  string `json:"username" binding:"required"`
		Password  string `json:"password" binding:"required"`
		Captcha   string `json:"captcha" binding:"required"`
		CaptchaId string `json:"captcha_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "请求参数错误")
		return
	}

	// 验证验证码
	if !VerifyCaptcha(req.CaptchaId, req.Captcha) {
		utils.BadRequest(c, "验证码错误")
		return
	}

	// 用户验证逻辑...
}
