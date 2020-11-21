package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/seanlan/wool/core"
)

func InitWebSocketRouter(router *gin.RouterGroup) {
	core.InitHub()
	wsRouter := router.Group("ws")
	wsRouter.GET("connect", core.CreateWebSocket)
}
