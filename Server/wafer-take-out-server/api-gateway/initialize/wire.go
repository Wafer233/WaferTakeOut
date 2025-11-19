//go:build wireinject

package initialize

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/internal/interfaces/rest"
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
		rpc.NewDishServiceClient,
		rpc.NewSetmealServiceClient,
		rpc.NewUserServiceClient,
		rpc.NewShopServiceClient,

		rpc.NewEmployeeService,
		rpc.NewCategoryService,
		rpc.NewDishService,
		rpc.NewSetMealService,
		rpc.NewUserService,
		rpc.NewShopService,

		rest.NewEmployeeHandler,
		rest.NewCategoryHandler,
		rest.NewDishHandler,
		rest.NewSetMealHandler,
		rest.NewCommonHandler,
		rest.NewUserHandler,
		rest.NewShopHandler,

		NewRouter,
	)

	return nil, nil
}
