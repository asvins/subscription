package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/asvins/common_io"
	"github.com/asvins/utils/config"
)

var (
	ServerConfig *Config = new(Config)
	producer     *common_io.Producer
	consumer     *common_io.Consumer
)

func init() {
	fmt.Println("[INFO] Initializing server")
	err := config.Load("subscription_config.gcfg", ServerConfig)
	if err != nil {
		log.Fatal(err)
	}

	setupCommonIo()
	fmt.Println("[INFO] Initialization Done!")
}

func main() {
	fmt.Println("[INFO] Server running on port:", ServerConfig.Server.Port)
	r := DefRoutes()
	http.ListenAndServe(":"+ServerConfig.Server.Port, r)
}
