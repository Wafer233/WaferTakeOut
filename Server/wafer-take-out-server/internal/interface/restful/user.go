package restful

import (
	categoryHandler "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/handler/category"
	dishHandler "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/handler/dish"
	setmealHandler "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/handler/setmeal"
	shopHandler "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/handler/shop"
	"github.com/gin-gonic/gin"
)

func NewUserRouter(
	r *gin.Engine,
	h1 *categoryHandler.CategoryHandler,
	h2 *setmealHandler.SetMealHandler,
	h3 *shopHandler.ShopHandler,
	h4 *dishHandler.DishHandler,
) *gin.Engine {

	category := r.Group("/user/category")
	category.GET("list", h1.GetCategoriesTyped)

	//setmeal := r.Group("/user/setmeal")
	//setmeal.GET("list", h2.GetById)

	shop := r.Group("/user/shop")
	shop.GET("status", h3.GetStatus)

	dish := r.Group("/user/dish")
	dish.GET("list", h4.GetDishesCategory)

	return r

}
