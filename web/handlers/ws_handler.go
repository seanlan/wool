package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/seanlan/packages/db"
	"github.com/seanlan/packages/logging"
	"github.com/seanlan/wool/core"
	"github.com/seanlan/wool/models"
	"github.com/seanlan/wool/utils"
	"net/http"
)

func CreateWebSocket(c *gin.Context) {
	webSocketHandler(c)
}

var wsUpGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func webSocketHandler(c *gin.Context) {
	var token, appKey, tag string
	token = c.Query("token")
	appKey = c.Query("appkey")
	var application models.ImApplication
	if err := db.DB.Where("appkey like ?", appKey).First(&application).Error; err != nil {
		c.String(403, "未登录或非法访问", "")
		c.Abort()
		return
	}
	j := utils.NewJWT(application.Appsecret)
	tag, err := j.ParseToken(token)
	if err != nil {
		logging.Logger.Debugf("token解析失败: %s", token)
		c.String(403, "未登录或非法访问", "")
		c.Abort()
		return
	}
	// 建立websocket连接
	conn, err := wsUpGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logging.Logger.Errorf("Failed to set websocket upgrade: %#v", err)
		return
	}
	core.NewWSClient(appKey, tag, conn)
}
