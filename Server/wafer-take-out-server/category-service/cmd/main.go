package cmd

import (
	"fmt"
	"net"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/category-service/internal/application"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/category-service/internal/infrastructure/database"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/category-service/internal/infrastructure/persistence"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/category-service/internal/interfaces/rpc"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/employee-service/proto"
	"google.golang.org/grpc"
)

func main() {

	db, _ := database.NewMysqlDatabase()
	repo := persistence.NewDefaultCategoryRepository(db)
	svc := application.NewCategoryService(repo)
	handler := rpc.NewCategoryHandler(svc)

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		panic(err)
	}
	fmt.Println("Employee gRPC service listening on localhost:50051")

	server := grpc.NewServer()
	proto.RegisterEmployeeServiceServer(server, handler)

	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}
