package main

import (
	"fmt"
	"net"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/shoppingcart-service/internal/application"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/shoppingcart-service/internal/infrastructure/database"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/shoppingcart-service/internal/infrastructure/persistence"
	rpcClient "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/shoppingcart-service/internal/infrastructure/rpc"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/shoppingcart-service/internal/interfaces/rpc"
	shoppingcartpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/shoppingcart-service/proto"
	"google.golang.org/grpc"
)

func main() {

	db, _ := database.NewMysqlDatabase()
	dishClient := rpcClient.NewDishServiceClient()
	setMealClient := rpcClient.NewSetMealServiceClient()
	dishSvc := rpcClient.NewDishService(dishClient)
	setMealSvc := rpcClient.NewSetMealService(setMealClient)
	repo := persistence.NewDefaultShoppingCartRepository(db)
	svc := application.NewShoppingCartService(repo, setMealSvc, dishSvc)
	handler := rpc.NewShoppingCartHandler(svc)

	lis, err := net.Listen("tcp", "localhost:50057")
	if err != nil {
		panic(err)
	}
	fmt.Println("ShoppingCart gRPC service listening on localhost:50057")

	server := grpc.NewServer()
	shoppingcartpb.RegisterShoppingCartServiceServer(server, handler)

	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}
