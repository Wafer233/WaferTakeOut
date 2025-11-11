package main

import (
	"log"

	categoryApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/category"
	commonApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/common"
	employeeApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/employee"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/infrastructure/persistence"
	categoryImpl "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/infrastructure/persistence/category"
	commonImpl "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/infrastructure/persistence/common"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/infrastructure/persistence/employee"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful"
	categoryHandler "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/handler/category"
	commonHandler "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/handler/common"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/handler/employee"
)

func main() {

	db, _ := persistence.NewMysqlDatabase()
	repo := employeeImpl.NewEmployeeRepository(db)
	repo1 := categoryImpl.NewCategoryRepository(db)
	repo2 := commonImpl.NewCommonRepository(db)
	service := employeeApp.NewEmployeeService(repo)
	service1 := categoryApp.NewCategoryService(repo1)
	service2 := commonApp.NewCommonService(repo2)
	h := employeeHandler.NewEmployeeHandler(service)
	h1 := categoryHandler.NewCategoryHandler(service1)
	h2 := commonHandler.NewCommonHandler(service2)
	r := restful.NewRouter(h, h1, h2)

	err := r.Run(":8080")

	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
