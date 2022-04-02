package authorization

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/khoale193/csv/pkg/e"
	"github.com/khoale193/csv/pkg/setting"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data interface{}
		code := e.SUCCESS
		token := c.Request.Header.Get(e.Authorization)
		if token != setting.AppSetting.Secret {
			code = e.INVALID_PARAMS
		}
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
