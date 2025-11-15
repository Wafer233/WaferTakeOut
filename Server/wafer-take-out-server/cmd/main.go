package main

import (
	"log"

	"github.com/Wafer233/WaferTakeOut/Server/wafer-take-out-server/internal/initialize"
)

func main() {

	r, err := initialize.Init()
	if err != nil {
		log.Fatal("初始化失败")
	}

	err = r.Run(":8080")

	if err != nil {
		log.Fatalf("启动服务器失败")
	}
}
