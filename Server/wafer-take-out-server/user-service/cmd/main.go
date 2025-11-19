package main

import (
	"fmt"
	"net"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/user-service/internal/application"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/user-service/internal/infrastructure/database"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/user-service/internal/infrastructure/persistence"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/user-service/internal/interfaces/rpc"
	userpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/user-service/proto"
	"google.golang.org/grpc"
)

func main() {

	db, _ := database.NewMysqlDatabase()
	repo := persistence.NewDefaultUserRepository(db)
	svc := application.NewUserService(repo)
	handler := rpc.NewUserHandler(svc)

	lis, err := net.Listen("tcp", "localhost:50055")
	if err != nil {
		panic(err)
	}
	fmt.Println("User gRPC service listening on localhost:50055")

	server := grpc.NewServer()
	userpb.RegisterUserServiceServer(server, handler)

	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}
