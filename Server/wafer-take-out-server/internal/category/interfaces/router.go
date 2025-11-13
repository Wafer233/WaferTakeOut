package interfaces

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, h *CategoryHandler) *gin.Engine {

	category := r.Group("/admin/category")
	category.Use(middleware.JWTAuthMiddleware())
	category.POST("", h.Create)
	category.GET("page", h.ListPage)
	category.PUT("", h.Update)
	category.POST("status/:status", h.UpdateStatus)
	category.DELETE("", h.Delete)
	category.GET("list", h.ListByType)
	return r
}
