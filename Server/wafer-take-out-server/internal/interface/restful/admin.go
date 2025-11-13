package restful

import (
	interfaces5 "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/common/interfaces"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/dish/interfaces"
	interfaces2 "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/employee/interfaces"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/handler"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/middleware"
	interfaces3 "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/setmeal/interfaces"
	interfaces4 "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/shop/interfaces"
	"github.com/gin-gonic/gin"
)

func NewAdminRouter(
	r *gin.Engine,
	h *interfaces2.EmployeeHandler,
	h1 *handler.CategoryHandler,
	h2 *interfaces5.CommonHandler,
	h3 *interfaces.DishHandler,
	h4 *interfaces3.SetMealHandler,
	h5 *interfaces4.ShopHandler,
) *gin.Engine {

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

	category := r.Group("/admin/category")
	category.Use(middleware.JWTAuthMiddleware())
	category.POST("", h1.Create)
	category.GET("page", h1.Page)
	category.PUT("", h1.Update)
	category.POST("status/:status", h1.UpdateStatus)
	category.DELETE("", h1.Delete)
	category.GET("list", h1.ListByType)

	common := r.Group("/admin/common")
	common.Use(middleware.JWTAuthMiddleware())
	common.POST("upload", h2.Upload)

	setMeal := r.Group("/admin/setmeal")
	setMeal.Use(middleware.JWTAuthMiddleware())
	setMeal.PUT("", h4.Update)
	setMeal.GET("page", h4.Page)
	setMeal.POST("status/:status", h4.UpdateStatus)
	setMeal.DELETE("", h4.Delete)
	setMeal.POST("", h4.Create)
	setMeal.GET(":id", h4.GetById)

	shop := r.Group("/admin/shop")
	shop.Use(middleware.JWTAuthMiddleware())
	shop.PUT(":status", h5.Update)
	shop.GET("status", h5.Get)

	return r
}
