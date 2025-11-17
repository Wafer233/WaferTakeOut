//go:build wireinject

package initialize

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/interface/rest"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/persistence/rpc"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type Application struct {
	Router *gin.Engine
}

func Init() (*gin.Engine, error) {
	wire.Build(

		rpc.NewEmployeeServiceClient,
		rpc.NewCategoryServiceClient,

		rpc.NewEmployeeService,
		rpc.NewCategoryService,

		rest.NewEmployeeHandler,
		rest.NewCategoryHandler,

		NewRouter,
	)

	return nil, nil
}
