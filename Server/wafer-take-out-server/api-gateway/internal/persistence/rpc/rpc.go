package rpc

import (
	categorypb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/proto/category"
	dishpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/proto/dish"
	employeepb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/proto/employee"
	setmealpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/proto/setmeal"
	"google.golang.org/grpc"
)

func NewEmployeeServiceClient() employeepb.EmployeeServiceClient {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return employeepb.NewEmployeeServiceClient(conn)
}

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

func NewSetmealServiceClient() setmealpb.SetmealServiceClient {
	conn, err := grpc.Dial("localhost:50054", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return setmealpb.NewSetmealServiceClient(conn)
}
