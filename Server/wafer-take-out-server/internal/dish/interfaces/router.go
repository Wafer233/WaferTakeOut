package interfaces

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, h *DishHandler) *gin.Engine {

	admin := r.Group("/admin/dish")

	admin.Use(middleware.EmployeeAuthMiddleware())

	admin.PUT("", h.Update)
	admin.DELETE("", h.Delete)
	admin.POST("", h.Create)
	admin.GET(":id", h.GetById)
	admin.GET("list", h.ListByCategory)
	admin.GET("page", h.Page)
	admin.POST("status/:status", h.UpdateStatus)

	user := r.Group("/user/dish")
	user.Use(middleware.UserAuthMiddleware())

	user.GET("list", h.ListByCategoryIdFlavor)

	return r
}
