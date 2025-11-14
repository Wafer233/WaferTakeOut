package interfaces

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, h *ShopHandler) *gin.Engine {

	admin := r.Group("/admin/shop")

	admin.Use(middleware.JWTAuthMiddleware())
	admin.PUT(":status", h.Update)
	admin.GET("status", h.Get)

	user := r.Group("/user/shop")
	//user.Use(middleware.JWTAuthMiddleware())
	user.GET("status", h.Get)
	return r

}
