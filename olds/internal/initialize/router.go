package initialize

import (
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/addressbook/internal/interfaces"
	"github.com/gin-gonic/gin"

	cateInter "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/category/interfaces"
	commonInter "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/common/interfaces"
	dishInter "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/dish/interfaces"
	emplInter "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/employee/interfaces"
	orderInter "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/order/interfaces"
	setmInter "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/setmeal/interfaces"
	shopInter "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/shop/interfaces"
	cartInter "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/shopping_cart/interfaces"
	userInter "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/user/interfaces"
)

func NewRouter(
	hEmpl *emplInter.EmployeeHandler,
	hCate *cateInter.CategoryHandler,
	hCommon *commonInter.CommonHandler,
	hDish *dishInter.DishHandler,
	hSetm *setmInter.SetMealHandler,
	hShop *shopInter.ShopHandler,
	hUser *userInter.UserHandler,
	hCart *cartInter.ShoppingCartHandler,
	hAddr *interfaces.AddressHandler,
	hOrder *orderInter.OrderHandler,
) *gin.Engine {

	r := gin.Default()

	r = emplInter.NewRouter(r, hEmpl)
	r = cateInter.NewRouter(r, hCate)
	r = commonInter.NewRouter(r, hCommon)
	r = dishInter.NewRouter(r, hDish)
	r = setmInter.NewRouter(r, hSetm)
	r = shopInter.NewRouter(r, hShop)
	r = userInter.NewRouter(r, hUser)
	r = cartInter.NewRouter(r, hCart)
	r = interfaces.NewRouter(r, hAddr)
	r = orderInter.NewRouter(r, hOrder)

	return r
}
