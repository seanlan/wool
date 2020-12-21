package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/seanlan/wool/core"
	"github.com/seanlan/wool/web/handlers"
)

func InitWebSocketRouter(router *gin.RouterGroup) {
	core.InitHub()
	wsRouter := router.Group("ws")
	wsRouter.GET("connect", handlers.CreateWebSocket)
}
