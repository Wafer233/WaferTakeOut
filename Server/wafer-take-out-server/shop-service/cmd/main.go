package main

import (
	"fmt"
	"net"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/shop-service/internal/application"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/shop-service/internal/infrastructure/database"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/shop-service/internal/infrastructure/persistence"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/shop-service/internal/interfaces/rpc"
	shoppb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/shop-service/proto"
	"google.golang.org/grpc"
)

func main() {

	db, _ := database.NewRedisDatabase()
	repo := persistence.NewDefaultShopRepository(db)
	svc := application.NewShopService(repo)
	handler := rpc.NewShopHandler(svc)

	lis, err := net.Listen("tcp", "localhost:50056")
	if err != nil {
		panic(err)
	}
	fmt.Println("Shop gRPC service listening on localhost:50056")

	server := grpc.NewServer()
	shoppb.RegisterShopServiceServer(server, handler)

	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}
