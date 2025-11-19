package initialize

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/interface/rest"
	"github.com/gin-gonic/gin"
)

func NewRouter(
	emp *rest.EmployeeHandler,
	cat *rest.CategoryHandler,
	dis *rest.DishHandler,
	sem *rest.SetMealHandler,
	com *rest.CommonHandler,
	use *rest.UserHandler,
) *gin.Engine {

	r := gin.Default()

	rest.NewEmployeeRouter(r, emp)
	rest.NewCategoryRouter(r, cat)
	rest.NewDishRouter(r, dis)
	rest.NewSetmealRouter(r, sem)
	rest.NewCommonRouter(r, com)
	rest.NewUserRouter(r, use)

	return r

}
