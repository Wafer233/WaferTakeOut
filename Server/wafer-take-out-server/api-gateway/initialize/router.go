package initialize

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/interface/rest"
	"github.com/gin-gonic/gin"
)

func NewRouter(emp *rest.EmployeeHandler) *gin.Engine {

	r := gin.Default()

	rest.NewEmployeeRouter(r, emp)

	return r

}
