package main

import "github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/api-gateway/initialize"

func main() {

	r, err := initialize.Init()
	if err != nil {
		panic(err)
	}

	_ = r.Run(":8080")
}
