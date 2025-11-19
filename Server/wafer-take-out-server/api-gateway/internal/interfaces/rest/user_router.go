package rest

import "github.com/gin-gonic/gin"

func NewUserRouter(r *gin.Engine, h *UserHandler) *gin.Engine {
	user := r.Group("/user/user")

	user.POST("/login", h.WxLogin)

	return r
}
