package main

import (
	"log"

	categoryApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/category"
	commonApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/common"
	dishApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/dish"
	employeeApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/employee"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/infrastructure/persistence"
	categoryImpl "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/infrastructure/persistence/category"
	commonImpl "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/infrastructure/persistence/common"
	dishImpl "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/infrastructure/persistence/dish"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/infrastructure/persistence/employee"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful"
	categoryHandler "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/handler/category"
	commonHandler "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/handler/common"
	dishHandler "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/handler/dish"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/handler/employee"
)

func main() {

	db, _ := persistence.NewMysqlDatabase()
	repo := employeeImpl.NewEmployeeRepository(db)
	repo1 := categoryImpl.NewCategoryRepository(db)
	repo2 := commonImpl.NewCommonRepository(db)
	repo3 := dishImpl.NewDishRepository(db)
	svc := employeeApp.NewEmployeeService(repo)
	svc1 := categoryApp.NewCategoryService(repo1)
	svc2 := commonApp.NewCommonService(repo2)
	svc3 := dishApp.NewDishService(repo3, repo1)
	h := employeeHandler.NewEmployeeHandler(svc)
	h1 := categoryHandler.NewCategoryHandler(svc1)
	h2 := commonHandler.NewCommonHandler(svc2)
	h3 := dishHandler.NewDishHandler(svc3)
	r := restful.NewRouter(h, h1, h2, h3)

	err := r.Run(":8080")

	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
