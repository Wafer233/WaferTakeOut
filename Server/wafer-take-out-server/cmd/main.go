package main

import (
	"log"

	categoryApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/category"
	employeeApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/employee"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/infrastructure/persistence"
	categoryImpl "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/infrastructure/persistence/category"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/infrastructure/persistence/employee"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful"
	categoryHandler "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/handler/category"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/handler/employee"
)

func main() {

	db, _ := persistence.NewMysqlDatabase()
	repo := employeeImpl.NewEmployeeRepository(db)
	repo1 := categoryImpl.NewCategoryRepository(db)
	service := employeeApp.NewEmployeeService(repo)
	service1 := categoryApp.NewCategoryService(repo1)
	h := employeeHandler.NewEmployeeHandler(service)
	h1 := categoryHandler.NewCategoryHandler(service1)
	r := restful.NewRouter(h, h1)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
