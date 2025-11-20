package main

import (
	"fmt"
	"net"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/order-service/internal/application"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/order-service/internal/domain"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/order-service/internal/infrastructure/database"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/order-service/internal/infrastructure/persistence"
	rpcClient "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/order-service/internal/infrastructure/rpc"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/order-service/internal/interfaces/rpc"
	orderpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/order-service/proto"

	"google.golang.org/grpc"
)

func main() {

	db, _ := database.NewMysqlDatabase()
	repo := persistence.NewDefaultOrderRepository(db)

	addr := rpcClient.NewAddressBookServiceClient()
	addrSvc := rpcClient.NewAddressBookService(addr)

	cart := rpcClient.NewShoppingCartServiceClient()
	cartSvc := rpcClient.NewShoppingCartService(cart)

	domainSvc := domain.NewOrderDomainService()

	svc := application.NewOrderService(repo, domainSvc, addrSvc, cartSvc)
	handler := rpc.NewOrderHandler(svc)

	lis, err := net.Listen("tcp", "localhost:50059")
	if err != nil {
		panic(err)
	}
	fmt.Println("Order gRPC service listening on localhost:50059")

	server := grpc.NewServer()
	orderpb.RegisterOrderServiceServer(server, handler)

	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}
