package interfaces

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, h *SetMealHandler) *gin.Engine {

	setMeal := r.Group("/admin/setmeal")
	setMeal.Use(middleware.JWTAuthMiddleware())
	setMeal.PUT("", h.Update)
	setMeal.GET("page", h.Page)
	setMeal.POST("status/:status", h.UpdateStatus)
	setMeal.DELETE("", h.Delete)
	setMeal.POST("", h.Create)
	setMeal.GET(":id", h.GetById)

	return r
}
