package rpc

import (
	categorypb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/setmeal-service/proto/category"
	dishpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/setmeal-service/proto/dish"
	"google.golang.org/grpc"
)

func NewCategoryServiceClient() categorypb.CategoryServiceClient {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return categorypb.NewCategoryServiceClient(conn)
}

func NewDishServiceClient() dishpb.DishServiceClient {
	conn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return dishpb.NewDishServiceClient(conn)
}
