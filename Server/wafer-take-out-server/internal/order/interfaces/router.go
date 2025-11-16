package interfaces

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, h *OrderHandler) *gin.Engine {

	user := r.Group("/user/order")
	user.Use(middleware.UserAuthMiddleware())

	user.POST("submit", h.Submit)
	user.PUT("/payment", h.Payment)
	return r
}
