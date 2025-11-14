package interfaces

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, h *EmployeeHandler) *gin.Engine {
	unprotected := r.Group("/admin/employee")
	unprotected.POST("/login", h.Login)
	unprotected.POST("/logout", h.Logout)

	employee := r.Group("/admin/employee")
	employee.Use(middleware.EmployeeAuthMiddleware())
	employee.POST("", h.Create)
	employee.GET("/page", h.ListPage)
	employee.POST("/status/:status", h.UpdateStatus)
	employee.GET("/:id", h.List)
	employee.PUT("", h.Update)
	employee.PUT("editPassword", h.UpdatePassword)
	return r
}
