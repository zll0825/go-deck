package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go-deck/app/global"
	"go-deck/app/response"
	"go.uber.org/zap"
)

var store = base64Captcha.DefaultMemStore

const (
	CaptchaImgHeight = 80
	CaptchaImgWidth = 240
	CaptchaImgLength = 4
	CaptchaMaxSkew = 0.7
	CaptchaDotCount = 80
)

// @Tags Base
// @Summary 生成验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证码获取成功"}"
// @Router /base/captcha [get]
func Captcha(c *gin.Context) {
	//字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(CaptchaImgHeight,CaptchaImgWidth,CaptchaImgLength, CaptchaMaxSkew,
		CaptchaDotCount)
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		global.Logger.Error("验证码获取失败!", zap.Any("err", err))
		response.FailWithMessage(c,"验证码获取失败")
	} else {
		response.OkWithDetailed(c, response.Captcha{
			CaptchaId: id,
			PicPath:   b64s,
		}, "验证码获取成功")
	}
}
