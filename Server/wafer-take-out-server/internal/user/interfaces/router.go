package interfaces

import "github.com/gin-gonic/gin"

func NewRouter(r *gin.Engine, h *UserHandler) *gin.Engine {
	user := r.Group("/user/user")

	user.POST("/login", h.WxLogin)

	return r
}
