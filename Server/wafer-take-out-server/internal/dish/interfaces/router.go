package interfaces

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, h *DishHandler) *gin.Engine {

	dish := r.Group("/admin/dish")

	dish.Use(middleware.JWTAuthMiddleware())
	dish.PUT("", h.Update)
	dish.DELETE("", h.Delete)
	dish.POST("", h.Create)
	dish.GET(":id", h.GetById)
	dish.GET("list", h.ListByCategory)
	dish.GET("page", h.Page)
	dish.POST("status/:status", h.UpdateStatus)
	return r
}
