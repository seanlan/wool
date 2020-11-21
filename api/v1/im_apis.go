package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/seanlan/packages/db"
	"github.com/seanlan/packages/router"
	"github.com/seanlan/wool/models"
	"github.com/seanlan/wool/utils"
)

// 获取IM Token
func GetIMToken(c *gin.Context) {
	var req struct {
		AppKey string `form:"appkey" json:"appkey" binding:"required"`
		UID    string `form:"uid" json:"uid" binding:"required"`
		Nonce  string `form:"nonce" json:"nonce" binding:"required"`
		Sign   string `form:"sign" json:"sign" binding:"required"`
	}
	err := router.RequestParser(&req, c)
	if err != nil {
		return
	}
	var application models.Application
	result := db.DB.Where("appkey like ?", req.AppKey).First(&application)
	if result.Error != nil {
		router.ErrorReturn(c, 1001, "app_key不存在")
		return
	}
	secretKey := application.Appsecret
	//验证请求合法性
	if utils.MapToUrlencoded(map[string]string{
		"appkey": req.AppKey,
		"uid":    req.UID,
		"nonce":  req.Nonce,
	}, secretKey) != req.Sign {
		router.ErrorReturn(c, 1002, "参数签名错误")
		return
	}
	imToken, _ := utils.MakeTimedToken(req.UID, secretKey, 3600*24*30)
	router.SuccessReturn(c, map[string]interface{}{"im_token": imToken, "uid": req.UID})
}
