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
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {

	db, _ := database.NewMysqlDatabase()
	rdb, _ := database.NewRedisDatabase()

	repo := emplRepo.NewEmployeeRepository(db)
	repo1 := cateRepo.NewCategoryRepository(db)
	repo3 := dishRepo.NewDishRepository(db)
	repo5 := setmRepo.NewSetMealRepository(db)
	cache := shopRepo.NewDefaultShopCache(rdb)
	svc := emplApp.NewEmployeeService(repo)
	svc1 := cateApp.NewCategoryService(repo1)
	svc3 := dishApp.NewDishService(repo3, repo1)
	svc4 := setmApp.NewSetMealService(repo5, repo1)
	svc5 := shopApp.NewShopService(cache)
	h := emplInter.NewEmployeeHandler(svc)
	h1 := cateInter.NewCategoryHandler(svc1)
	h2 := commonInter.NewCommonHandler()
	h3 := dishInter.NewDishHandler(svc3)
	h4 := setmInter.NewSetMealHandler(svc4)
	h5 := shopInter.NewShopHandler(svc5)

	r := gin.Default()
	r = emplInter.NewRouter(r, h)
	r = cateInter.NewRouter(r, h1)
	r = commonInter.NewRouter(r, h2)
	r = dishInter.NewRouter(r, h3)
	r = setmInter.NewRouter(r, h4)
	r = shopInter.NewRouter(r, h5)

	err := r.Run(":8080")

	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
