package rest

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/interfaces/rest/middleware"
	"github.com/gin-gonic/gin"
)

func NewShopRouter(r *gin.Engine, h *ShopHandler) *gin.Engine {

	admin := r.Group("/admin/shop")

	admin.Use(middleware.EmployeeAuthMiddleware())
	admin.PUT(":status", h.Update)
	admin.GET("status", h.Get)

	user := r.Group("/user/shop")

	user.GET("status", h.Get)
	return r

}
