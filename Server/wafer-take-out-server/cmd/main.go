package main

import (
	"log"

	commonApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/common"
	categoryApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/category/application"
	categoryImpl "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/category/infrastructure"
	dishApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/dish/application"
	dishImpl "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/dish/infrastructure"
	employeeApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/employee/application"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/employee/infrastructure"
	commonImpl "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/infrastructure/persistence/common"
	flavorImpl "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/infrastructure/persistence/flavor"
	setmealDishImpl "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/infrastructure/persistence/setmeal_dish"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful"
	categoryHandler "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/handler/category"
	commonHandler "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/handler/common"
	dishHandler "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/handler/dish"
	setmealHandler "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/handler/setmeal"
	shopHandler "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/handler/shop"
	setmealApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/setmeal/application"
	setmealImpl "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/setmeal/infrastructure"
	shopApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/shop/application"
	shopImpl "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/shop/infrastructure"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {

	db, _ := database.NewMysqlDatabase()
	rdb, _ := database.NewRedisDatabase()
	repo := infrastructure.employeeImpl.NewEmployeeRepository(db)
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

	r := gin.Default()
	r = restful.NewAdminRouter(r, h, h1, h2, h3, h4, h5)
	r = restful.NewUserRouter(r, h1, h4, h5, h3)

	err := r.Run(":8080")

	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
