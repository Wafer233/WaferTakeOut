//go:build wireinject

package initialize

import (
	"github.com/google/wire"

	"github.com/gin-gonic/gin"

	cateApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/category/application"
	cateRepo "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/category/infrastructure"
	cateInter "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/category/interfaces"

	commonInter "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/common/interfaces"

	dishApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/dish/application"
	dishRepo "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/dish/infrastructure"
	dishInter "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/dish/interfaces"

	emplApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/employee/application"
	emplRepo "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/employee/infrastructure"
	emplInter "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/employee/interfaces"

	setmApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/setmeal/application"
	setmRepo "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/setmeal/infrastructure"
	setmInter "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/setmeal/interfaces"

	shopApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/shop/application"
	shopRepo "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/shop/infrastructure"
	shopInter "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/shop/interfaces"

	cartApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/shopping_cart/application"
	cartRepo "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/shopping_cart/infrastructure"
	cartInter "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/shopping_cart/interfaces"

	userApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/user/application"
	userRepo "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/user/infrastructure"
	userInter "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/user/interfaces"

	addrApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/addressbook/application"
	addrRepo "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/addressbook/infrastructure"
	addrInter "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/addressbook/interfaces"

	orderApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/order/application"
	orderRepo "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/order/infrastructure"
	orderInter "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/order/interfaces"
)

type Application struct {
	Router *gin.Engine
}

func Init() (*gin.Engine, error) {
	wire.Build(

		// database providers
		NewMysqlDatabase,
		NewRedisDatabase,

		// repositories
		// 这里注意一下，我的New返回的是接口，但是default和cached两个都是接口，所以只能一个返回接口
		emplRepo.NewDefaultEmployeeRepository,
		cateRepo.NewDefaultCategoryRepository,
		dishRepo.NewDefaultDishRepository,
		dishRepo.NewCachedDishRepository,
		setmRepo.NewDefaultSetMealRepository,
		shopRepo.NewDefaultShopRepository,
		userRepo.NewDefaultUserRepository,
		cartRepo.NewDefaultShoppingCartRepository,
		addrRepo.NewDefaultAddressRepository,
		orderRepo.NewDefaultOrderRepository,

		// services
		emplApp.NewEmployeeService,
		cateApp.NewCategoryService,
		dishApp.NewDishService,
		setmApp.NewSetMealService,
		shopApp.NewShopService,
		userApp.NewUserService,
		cartApp.NewShoppingCartService,
		addrApp.NewAddressService,
		orderApp.NewOrderService,

		// handlers
		emplInter.NewEmployeeHandler,
		cateInter.NewCategoryHandler,
		commonInter.NewCommonHandler,
		dishInter.NewDishHandler,
		setmInter.NewSetMealHandler,
		shopInter.NewShopHandler,
		userInter.NewUserHandler,
		cartInter.NewShoppingCartHandler,
		addrInter.NewAddressHandler,
		orderInter.NewOrderHandler,

		// router
		NewRouter,
	)

	return nil, nil
}
