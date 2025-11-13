package interfaces

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewAdminRouter(r *gin.Engine, h *EmployeeHandler) *gin.Engine {
	unprotected := r.Group("/admin/employee")
	unprotected.POST("/login", h.Login)
	unprotected.POST("/logout", h.Logout)

	employee := r.Group("/admin/employee")
	employee.Use(middleware.JWTAuthMiddleware())
	employee.POST("", h.Create)
	employee.GET("/page", h.Page)
	employee.POST("/status/:status", h.UpdateStatus)
	employee.GET("/:id", h.GetById)
	employee.PUT("", h.Update)
	return r
}
