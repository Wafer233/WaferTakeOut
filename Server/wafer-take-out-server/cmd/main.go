package main

import (
	"log"

	categoryApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/category"
	commonApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/common"
	dishApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/dish"
	employeeApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/employee"
	setmealApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/setmeal"
	shopApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/shop"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/infrastructure/persistence"
	categoryImpl "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/infrastructure/persistence/category"
	commonImpl "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/infrastructure/persistence/common"
	dishImpl "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/infrastructure/persistence/dish"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/infrastructure/persistence/employee"
	flavorImpl "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/infrastructure/persistence/flavor"
	setmealImpl "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/infrastructure/persistence/setmeal"
	setmealDishImpl "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/infrastructure/persistence/setmeal_dish"
	shopImpl "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/infrastructure/persistence/shop"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful"
	categoryHandler "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/handler/category"
	commonHandler "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/handler/common"
	dishHandler "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/handler/dish"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/handler/employee"
	setmealHandler "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/handler/setmeal"
	shopHandler "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/handler/shop"
)

func main() {

	db, _ := persistence.NewMysqlDatabase()
	rdb, _ := persistence.NewRedisDatabase()
	repo := employeeImpl.NewEmployeeRepository(db)
	repo1 := categoryImpl.NewCategoryRepository(db)
	repo2 := commonImpl.NewCommonRepository(db)
	repo3 := dishImpl.NewDishRepository(db)
	repo4 := flavorImpl.NewFlavorRepository(db)
	repo5 := setmealImpl.NewSetMealRepository(db)
	repo6 := setmealDishImpl.NewSetMealDishRepository(db)
	cache := shopImpl.NewDefaultShopCache(rdb)
	svc := employeeApp.NewEmployeeService(repo)
	svc1 := categoryApp.NewCategoryService(repo1)
	svc2 := commonApp.NewCommonService(repo2)
	svc3 := dishApp.NewDishService(repo3, repo1, repo4)
	svc4 := setmealApp.NewSetMealService(repo5, repo6, repo1)
	svc5 := shopApp.NewShopService(cache)
	h := employeeHandler.NewEmployeeHandler(svc)
	h1 := categoryHandler.NewCategoryHandler(svc1)
	h2 := commonHandler.NewCommonHandler(svc2)
	h3 := dishHandler.NewDishHandler(svc3)
	h4 := setmealHandler.NewSetMealHandler(svc4)
	h5 := shopHandler.NewShopHandler(svc5)
	r := restful.NewRouter(h, h1, h2, h3, h4, h5)

	err := r.Run(":8080")

	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
