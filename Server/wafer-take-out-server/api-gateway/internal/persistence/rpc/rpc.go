package rpc

import (
	employeepb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/proto/employee"
	"google.golang.org/grpc"
)

func NewEmployeeServiceClient() employeepb.EmployeeServiceClient {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return employeepb.NewEmployeeServiceClient(conn)
}
