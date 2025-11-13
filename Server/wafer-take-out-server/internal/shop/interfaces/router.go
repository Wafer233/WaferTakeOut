package interfaces

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, h5 *ShopHandler) *gin.Engine {

	shop := r.Group("/admin/shop")

	shop.Use(middleware.JWTAuthMiddleware())
	shop.PUT(":status", h5.Update)
	shop.GET("status", h5.Get)
	return r

}
