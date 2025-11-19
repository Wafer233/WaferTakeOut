package main

import (
	"fmt"
	"net"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/setmeal-service/internal/application"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/setmeal-service/internal/infrastructure/database"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/setmeal-service/internal/infrastructure/persistence"
	rpcClient "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/setmeal-service/internal/infrastructure/rpc"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/setmeal-service/internal/interfaces/rpc"
	setmealpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/setmeal-service/proto"
	"google.golang.org/grpc"
)

func main() {

	db, _ := database.NewMysqlDatabase()
	repo := persistence.NewDefaultSetMealRepository(db)
	cateClient := rpcClient.NewCategoryServiceClient()
	cateSvc := rpcClient.NewCategoryService(cateClient)
	dishClient := rpcClient.NewDishServiceClient()
	dishSvc := rpcClient.NewDishService(dishClient)

	svc := application.NewSetMealService(repo, cateSvc, dishSvc)
	handler := rpc.NewSetMealHandler(svc)

	lis, err := net.Listen("tcp", "localhost:50054")
	if err != nil {
		panic(err)
	}
	fmt.Println("SetMeal gRPC service listening on localhost:50054")

	server := grpc.NewServer()
	setmealpb.RegisterSetmealServiceServer(server, handler)

	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}
