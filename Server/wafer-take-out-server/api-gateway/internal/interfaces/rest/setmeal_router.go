package rest

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/interfaces/rest/middleware"
	"github.com/gin-gonic/gin"
)

func NewSetmealRouter(r *gin.Engine, h *SetMealHandler) *gin.Engine {

	admin := r.Group("/admin/setmeal")
	admin.Use(middleware.EmployeeAuthMiddleware())

	admin.PUT("", h.Update)
	admin.GET("page", h.ListPage)
	admin.POST("status/:status", h.UpdateStatus)
	admin.DELETE("", h.Delete)
	admin.POST("", h.Create)
	admin.GET(":id", h.ListById)

	user := r.Group("/user/setmeal")
	user.Use(middleware.UserAuthMiddleware())

	user.GET("list", h.ListByCategoryId)
	user.GET("dish/:id", h.ListDishById)
	return r
}
