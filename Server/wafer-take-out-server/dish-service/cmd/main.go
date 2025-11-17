package main

import (
	"fmt"
	"net"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/dish-service/internal/application"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/dish-service/internal/infrastructure/database"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/dish-service/internal/infrastructure/persistence"
	rpcClient "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/dish-service/internal/infrastructure/rpc"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/dish-service/internal/interfaces/rpc"
	dishpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/dish-service/proto"
	"google.golang.org/grpc"
)

func main() {

	db, _ := database.NewMysqlDatabase()
	rdb, _ := database.NewRedisDatabase()
	cache := persistence.NewDefaultDishRepository(db)
	repo := persistence.NewCachedDishRepository(cache, rdb)

	client := rpcClient.NewCategoryServiceClient()
	cateSvc := rpcClient.NewCategoryService(client)

	svc := application.NewDishService(repo, cateSvc)
	handler := rpc.NewDishHandler(svc)

	lis, err := net.Listen("tcp", "localhost:50053")
	if err != nil {
		panic(err)
	}
	fmt.Println("Dish gRPC service listening on localhost:50053")

	server := grpc.NewServer()
	dishpb.RegisterDishServiceServer(server, handler)

	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}
