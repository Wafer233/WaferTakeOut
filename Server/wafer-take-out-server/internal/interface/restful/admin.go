package restful

import (
	categoryHandler "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/handler/category"
	commonHandler "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/handler/common"
	dishHandler "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/handler/dish"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/handler/employee"
	setmealHandler "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/handler/setmeal"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(h *employeeHandler.EmployeeHandler,
	h1 *categoryHandler.CategoryHandler,
	h2 *commonHandler.CommonHandler,
	h3 *dishHandler.DishHandler,
	h4 *setmealHandler.SetMealHandler,
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
	category.DELETE("", h1.DeleteCategory)
	category.GET("list", h1.GetCategoriesTyped)

	common := r.Group("/admin/common")
	common.Use(middleware.JWTAuthMiddleware())
	common.POST("upload", h2.Upload)

	dish := r.Group("/admin/dish")
	dish.Use(middleware.JWTAuthMiddleware())
	dish.PUT("", h3.EditDish)
	dish.DELETE("", h3.DeleteDishes)
	dish.POST("", h3.AddDish)
	dish.GET(":id", h3.GetDishId)
	dish.GET("list", h3.GetDishesCategory)
	dish.GET("page", h3.GetDishesPaged)
	dish.POST("status/:status", h3.EditDishStatus)

	setMeal := r.Group("/admin/setmeal")
	setMeal.Use(middleware.JWTAuthMiddleware())
	setMeal.GET("page", h4.GetSetMealsPaged)
	setMeal.POST("status/:status", h4.EditSetMealStatus)
	setMeal.DELETE("", h4.DeleteSetMeal)
	setMeal.POST("", h4.AddSetMeal)

	return r
}
