package restful

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/handler"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(h *handler.EmployeeHandler) *gin.Engine {
	r := gin.Default()

	admin := r.Group("/admin/employee")
	admin.POST("/login", h.Login)
	admin.POST("/logout", h.Logout)

	protected := r.Group("/admin/employee")
	protected.Use(middleware.JWTAuthMiddleware())
	protected.POST("", h.AddEmployee)
	protected.GET("/page", h.Page)
	protected.POST("/status/:status", h.StatusFlip)

	return r
}
