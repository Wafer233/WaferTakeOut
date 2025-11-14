package main

import (
	"log"

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
	userRepo "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/user/infrastructure"
	userInter "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/user/interfaces"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/database"
	"github.com/gin-gonic/gin"

	userApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/user/application"
)

func main() {

	db, _ := database.NewMysqlDatabase()
	rdb, _ := database.NewRedisDatabase()

	repo := emplRepo.NewDefaultEmployeeRepository(db)
	repo1 := cateRepo.NewDefaultCategoryRepository(db)
	repo3 := dishRepo.NewDefaultDishRepository(db)
	repo5 := setmRepo.NewDefaultSetMealRepository(db)
	repo6 := shopRepo.NewDefaultShopRepository(rdb)
	repo7 := userRepo.NewDefaultUserRepository(db)
	repo8 := dishRepo.NewCachedDishRepository(repo3, rdb)
	repo9 := cartRepo.NewDefaultShoppingCartRepository(db)

	svc := emplApp.NewEmployeeService(repo)
	svc1 := cateApp.NewCategoryService(repo1)
	svc3 := dishApp.NewDishService(repo8, repo1)
	svc4 := setmApp.NewSetMealService(repo5, repo1, repo3)
	svc5 := shopApp.NewShopService(repo6)
	svc6 := userApp.NewUserService(repo7)
	svc7 := cartApp.NewShoppingCartService(repo9)

	h := emplInter.NewEmployeeHandler(svc)
	h1 := cateInter.NewCategoryHandler(svc1)
	h2 := commonInter.NewCommonHandler()
	h3 := dishInter.NewDishHandler(svc3)
	h4 := setmInter.NewSetMealHandler(svc4)
	h5 := shopInter.NewShopHandler(svc5)
	h6 := userInter.NewUserHandler(svc6)
	h7 := cartInter.NewShoppingCartHandler(svc7)

	r := gin.Default()
	r = emplInter.NewRouter(r, h)
	r = cateInter.NewRouter(r, h1)
	r = commonInter.NewRouter(r, h2)
	r = dishInter.NewRouter(r, h3)
	r = setmInter.NewRouter(r, h4)
	r = shopInter.NewRouter(r, h5)
	r = userInter.NewRouter(r, h6)
	r = cartInter.NewRouter(r, h7)

	err := r.Run(":8080")

	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
