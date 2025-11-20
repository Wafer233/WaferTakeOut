package main

import (
	"fmt"
	"net"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/addressbook-service/internal/application"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/addressbook-service/internal/infrastructure/database"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/addressbook-service/internal/infrastructure/persistence"
	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/addressbook-service/internal/interfaces/rpc"
	addressbookpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/addressbook-service/proto"
	"google.golang.org/grpc"
)

func main() {

	db, _ := database.NewMysqlDatabase()
	repo := persistence.NewDefaultAddressRepository(db)
	svc := application.NewAddressService(repo)
	handler := rpc.NewAddressBookHandler(svc)

	lis, err := net.Listen("tcp", "localhost:50058")
	if err != nil {
		panic(err)
	}
	fmt.Println("AddressBook" +
		" gRPC service listening on localhost:50058")

	server := grpc.NewServer()
	addressbookpb.RegisterAddressBookServiceServer(server, handler)

	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}
