package rpc

import (
	addressbookpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/order-service/proto/addressbook"
	shoppingcartpb "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/order-service/proto/shoppingcart"
	"google.golang.org/grpc"
)

func NewAddressBookServiceClient() addressbookpb.AddressBookServiceClient {
	conn, err := grpc.Dial("localhost:50058", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return addressbookpb.NewAddressBookServiceClient(conn)
}

func NewShoppingCartServiceClient() shoppingcartpb.ShoppingCartServiceClient {
	conn, err := grpc.Dial("localhost:50057", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return shoppingcartpb.NewShoppingCartServiceClient(conn)
}
