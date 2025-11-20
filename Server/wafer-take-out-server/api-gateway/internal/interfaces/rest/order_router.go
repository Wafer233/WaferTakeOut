package rest

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/interfaces/rest/middleware"
	"github.com/gin-gonic/gin"
)

func NewOrderRouter(r *gin.Engine, h *OrderHandler) *gin.Engine {

	user := r.Group("/user/order")
	user.Use(middleware.UserAuthMiddleware())

	user.POST("submit", h.Submit)
	user.PUT("/payment", h.Payment)
	user.GET("/historyOrders", h.ListPage)
	user.GET("/orderDetail/:id", h.GetOrder)
	user.PUT("/cancel/:id", h.UserCancel)
	user.POST("repetition/:id", h.CreateSame)

	admin := r.Group("/admin/order")
	admin.Use(middleware.EmployeeAuthMiddleware())
	admin.GET("/conditionSearch", h.ListAdminPage)
	admin.GET("/statistics", h.GetStatistics)
	admin.GET("/details/:id", h.GetOrder)
	admin.PUT("/confirm", h.Confirm)
	admin.PUT("/rejection", h.Rejection)
	admin.PUT("/cancel", h.Cancel)
	admin.PUT("/delivery/:id", h.Delivery)
	admin.PUT("/complete/:id", h.Complete)

	return r
}
