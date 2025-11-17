package middleware

import (
	"net/http"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/pkg/jwt"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/pkg/result"
	"github.com/gin-gonic/gin"
)

// EmployeeAuthMiddleware
// 这里就是两个版本 版本1 后端生成通过上下文生成一个cookie，每一次需要权限校验的解析cookie
func EmployeeAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 从cookie中获取token
		token, err := c.Cookie("token")
		if err != nil || token == "" {
			c.JSON(http.StatusUnauthorized, result.Error("未授权"))
			c.Abort()
			return
		}

		claims, err := jwt.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, result.Error("未授权"))
			c.Abort()
			return
		}

		c.Set("CurID", claims.ID)
		c.Next()
	}
}

// UserAuthMiddleware
// 一种是登陆时生成token发送给前端，前端在每一次请求中都带上对应的header
// 比如这里就是authentication
func UserAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("authentication")
		if token == "" {
			c.JSON(http.StatusUnauthorized, result.Error("未授权"))
			c.Abort()
			return
		}

		claim, err := jwt.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, result.Error("未授权"))
			c.Abort()
			return
		}

		c.Set("CurID", claim.ID)
		c.Next()
	}
}
