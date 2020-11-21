package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/seanlan/wool/api/v1"
)

func InitIMApiRouter(router *gin.RouterGroup) {
	r := router.Group("im")
	r.POST("get_token", v1.GetIMToken)
}
