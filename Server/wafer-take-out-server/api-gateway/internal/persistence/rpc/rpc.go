package rpc

import (
	addressbookpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/proto/addressbook"
	categorypb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/proto/category"
	dishpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/proto/dish"
	employeepb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/proto/employee"
	orderpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/proto/order"
	setmealpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/proto/setmeal"
	shoppb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/proto/shop"
	shoppingcartpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/proto/shoppingcart"
	userpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/proto/user"
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

func NewUserServiceClient() userpb.UserServiceClient {
	conn, err := grpc.Dial("localhost:50055", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return userpb.NewUserServiceClient(conn)
}

func NewShopServiceClient() shoppb.ShopServiceClient {
	conn, err := grpc.Dial("localhost:50056", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return shoppb.NewShopServiceClient(conn)
}

func NewShoppingCartServiceClient() shoppingcartpb.ShoppingCartServiceClient {
	conn, err := grpc.Dial("localhost:50057", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return shoppingcartpb.NewShoppingCartServiceClient(conn)
}

func NewAddressBookServiceClient() addressbookpb.AddressBookServiceClient {
	conn, err := grpc.Dial("localhost:50058", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return addressbookpb.NewAddressBookServiceClient(conn)
}

func NewOrderServiceClient() orderpb.OrderServiceClient {
	conn, err := grpc.Dial("localhost:50059", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return orderpb.NewOrderServiceClient(conn)
}
