package interfaces

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, h *SetMealHandler) *gin.Engine {

	admin := r.Group("/admin/setmeal")
	admin.Use(middleware.JWTAuthMiddleware())
	admin.PUT("", h.Update)
	admin.GET("page", h.ListPage)
	admin.POST("status/:status", h.UpdateStatus)
	admin.DELETE("", h.Delete)
	admin.POST("", h.Create)
	admin.GET(":id", h.ListById)

	user := r.Group("/user/setmeal")
	//user.Use(middleware.JWTAuthMiddleware())
	user.GET("list", h.ListByCategoryId)
	user.GET("dish/:id", h.ListDishById)
	return r
}
