package restful

import (
	categoryHandler "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/handler/category"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/handler/employee"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(h *employeeHandler.EmployeeHandler,
	h1 *categoryHandler.CategoryHandler,
) *gin.Engine {
	r := gin.Default()

	employee := r.Group("/admin/employee")
	employee.POST("/login", h.Login)
	employee.POST("/logout", h.Logout)

	protected := r.Group("/admin/employee")
	protected.Use(middleware.JWTAuthMiddleware())
	protected.POST("", h.AddEmployee)
	protected.GET("/page", h.Page)
	protected.POST("/status/:status", h.StatusFlip)
	protected.GET("/:id", h.GetEmployee)
	protected.PUT("", h.EditEmployee)

	category := r.Group("/admin/category")
	category.Use(middleware.JWTAuthMiddleware())
	category.POST("", h1.AddCategory)
	category.GET("page", h1.GetCategories)
	category.PUT("", h1.EditCategory)
	category.POST("status/:status", h1.EditCategoryStatus)

	return r
}
