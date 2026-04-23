package main

import (
	"fmt"
	"net/http"

	"github.com/yury-nazarov/purple-order-api/configs"
	"github.com/yury-nazarov/purple-order-api/pkg/db"
)

func main() {
	config := configs.LoadConfig()
	_ = db.New(config.Db.Dsn)

	server := http.Server{
		Addr: ":8082",
		// Handler: router
	}

	fmt.Println("HTTP Server is running on 8082")
	server.ListenAndServe()
}
