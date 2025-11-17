package interfaces

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, h *ShoppingCartHandler) *gin.Engine {

	// 要小心pattern会不会多一个空格这样
	user := r.Group("/user/shoppingCart/")
	user.Use(middleware.UserAuthMiddleware())

	user.POST("/add", h.Add)
	user.GET("/list", h.List)
	user.POST("/sub", h.Sub)
	user.DELETE("/clean", h.Delete)

	return r
}
