package initialize

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/interfaces/rest"
	"github.com/gin-gonic/gin"
)

func NewRouter(
	emp *rest.EmployeeHandler,
	cat *rest.CategoryHandler,
	dis *rest.DishHandler,
	sem *rest.SetMealHandler,
	com *rest.CommonHandler,
	use *rest.UserHandler,
	sho *rest.ShopHandler,
) *gin.Engine {

	r := gin.Default()

	rest.NewEmployeeRouter(r, emp)
	rest.NewCategoryRouter(r, cat)
	rest.NewDishRouter(r, dis)
	rest.NewSetmealRouter(r, sem)
	rest.NewCommonRouter(r, com)
	rest.NewUserRouter(r, use)
	rest.NewShopRouter(r, sho)

	return r

}
