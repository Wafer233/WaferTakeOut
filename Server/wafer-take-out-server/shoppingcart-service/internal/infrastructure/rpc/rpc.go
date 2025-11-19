package rpc

import (
	dishpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/shoppingcart-service/proto/dish"
	setmealpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/shoppingcart-service/proto/setmeal"
	"google.golang.org/grpc"
)

func NewDishServiceClient() dishpb.DishServiceClient {

	conn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return dishpb.NewDishServiceClient(conn)
}

func NewSetMealServiceClient() setmealpb.SetmealServiceClient {
	conn, err := grpc.Dial("localhost:50054", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return setmealpb.NewSetmealServiceClient(conn)
}
