package interfaces

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, h *CategoryHandler) *gin.Engine {

	admin := r.Group("/admin/category")
	admin.Use(middleware.JWTAuthMiddleware())
	admin.POST("", h.Create)
	admin.GET("page", h.ListPage)
	admin.PUT("", h.Update)
	admin.POST("status/:status", h.UpdateStatus)
	admin.DELETE("", h.Delete)
	admin.GET("list", h.ListByType)

	user := r.Group("/user/category")
	//user.Use(middleware.JWTAuthMiddleware())
	user.GET("list", h.ListByType)
	return r
}
