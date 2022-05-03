package main

import (
	"github.com/RadisaMilovcevic/bsep_xws/microservices_demo/shipping_service/startup"
	cfg "github.com/RadisaMilovcevic/bsep_xws/microservices_demo/shipping_service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
