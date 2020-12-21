package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/seanlan/packages/db"
	"github.com/seanlan/packages/router"
	"github.com/seanlan/wool/models"
	"github.com/seanlan/wool/service"
	"github.com/seanlan/wool/utils"
	"github.com/seanlan/wool/utils/wool_sdk"
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
	var application models.ImApplication
	result := db.DB.Where("appkey like ?", req.AppKey).First(&application)
	if result.Error != nil {
		router.ErrorReturn(c, 1001, "app_key不存在")
		return
	}
	sdk := wool_sdk.WoolSDK{
		AppKey:    application.Appkey,
		AppSecret: application.Appsecret,
	}
	params := map[string]string{
		"appkey": req.AppKey,
		"uid":    req.UID,
		"nonce":  req.Nonce}
	//验证请求合法性
	if sdk.GetSign(params) != req.Sign {
		router.ErrorReturn(c, 1002, "参数签名错误")
		return
	}
	imToken, _ := utils.MakeTimedToken(req.UID, application.Appsecret, 3600*24*30)
	router.SuccessReturn(c, map[string]interface{}{"im_token": imToken, "uid": req.UID})
}

// 创建会话
func CreateSingleConversation(c *gin.Context) {
	var req struct {
		AppKey string `form:"appkey" json:"appkey" binding:"required"`
		From   string `form:"from" json:"from" binding:"required"`
		To     string `form:"to" json:"to" binding:"required"`
		Title  string `form:"title" json:"title" binding:"required"`
		Icon   string `form:"icon" json:"icon" binding:"required"`
		Nonce  string `form:"nonce" json:"nonce" binding:"required"`
		Sign   string `form:"sign" json:"sign" binding:"required"`
	}
	err := router.RequestParser(&req, c)
	if err != nil {
		return
	}
	var application models.ImApplication
	result := db.DB.Where("appkey like ?", req.AppKey).First(&application)
	if result.Error != nil {
		router.ErrorReturn(c, 1001, "app_key不存在")
		return
	}
	sdk := wool_sdk.WoolSDK{
		AppKey:    application.Appkey,
		AppSecret: application.Appsecret,
	}
	params := map[string]string{
		"appkey": req.AppKey,
		"from":   req.From,
		"to":     req.To,
		"title":  req.Title,
		"icon":   req.Icon,
		"nonce":  req.Nonce}
	//验证请求合法性
	if sdk.GetSign(params) != req.Sign {
		router.ErrorReturn(c, 1002, "参数签名错误")
		return
	}
	conversationID :=service.CreateSingleConversation(req.AppKey,req.From,req.To,req.Title,req.Icon)
	router.SuccessReturn(c, map[string]interface{}{"conversation_id": conversationID})
}
