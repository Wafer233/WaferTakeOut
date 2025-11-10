package main

import (
	"log"

	employeeApp "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/application/employee"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/infrastructure/persistence"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/infrastructure/persistence/employee"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/interface/restful/handler/employee"
)

func main() {

	db, _ := persistence.NewMysqlDatabase()
	repo := employeeImpl.NewEmployeeRepository(db)
	service := employeeApp.NewEmployeeService(repo)
	h := employeeHandler.NewEmployeeHandler(service)
	r := restful.NewRouter(h)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
