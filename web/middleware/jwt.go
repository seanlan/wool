package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/seanlan/packages/logging"
	"github.com/seanlan/wool/utils"
	"strings"
)

func JWTAuth(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token
		// 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中
		// 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		//var token string
		var err error
		var token string
		switch strings.ToUpper(c.Request.Method) {
		case "GET":
			token = c.Query("token")
			logging.Logger.Debugf("GET 请求: %v", token)
		case "POST":
			var req struct {
				Token string `json:"token" binding:"required"`
			}
			logging.Logger.Debugf("POST 请求")
			err = c.ShouldBind(&req)
			if err != nil {
				logging.Logger.Debugf("未登录或非法访问")
				c.String(403, "未登录或非法访问", "")
				c.Abort()
				return
			}
			token = req.Token
		}
		j := utils.NewJWT(secretKey)
		userID, err := j.ParseToken(token)
		if err != nil {
			logging.Logger.Debugf("token解析失败: %s", token)
			c.String(403, "未登录或非法访问", "")
			c.Abort()
			return
		}
		c.Set("userID", userID)
		c.Next()
	}
}
