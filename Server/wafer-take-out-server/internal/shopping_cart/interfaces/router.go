package interfaces

import "github.com/gin-gonic/gin"

func NewRouter(r *gin.Engine, h *ShoppingCartHandler) *gin.Engine {

	// 要小心pattern会不会多一个空格这样
	user := r.Group("/user/shoppingCart/")

	user.POST("/add", h.Create)

	return r
}
