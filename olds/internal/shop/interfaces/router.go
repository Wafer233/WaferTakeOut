package interfaces

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, h *ShopHandler) *gin.Engine {

	admin := r.Group("/admin/shop")

	admin.Use(middleware.EmployeeAuthMiddleware())
	admin.PUT(":status", h.Update)
	admin.GET("status", h.Get)

	user := r.Group("/user/shop")
	user.Use(middleware.UserAuthMiddleware())

	user.GET("status", h.Get)
	return r

}
