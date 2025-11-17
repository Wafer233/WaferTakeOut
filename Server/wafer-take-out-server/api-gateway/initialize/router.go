package initialize

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/interface/rest"
	"github.com/gin-gonic/gin"
)

func NewRouter(
	emp *rest.EmployeeHandler,
	cat *rest.CategoryHandler,
	dis *rest.DishHandler,
) *gin.Engine {

	r := gin.Default()

	rest.NewEmployeeRouter(r, emp)
	rest.NewCategoryRouter(r, cat)
	rest.NewDishRouter(r, dis)

	return r

}
