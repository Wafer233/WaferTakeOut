package rpc

import (
	categorypb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/dish-service/proto/category"
	"google.golang.org/grpc"
)

func NewCategoryServiceClient() categorypb.CategoryServiceClient {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return categorypb.NewCategoryServiceClient(conn)
}
