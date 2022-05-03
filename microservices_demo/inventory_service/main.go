package main

import (
	"github.com/RadisaMilovcevic/bsep_xws/microservices_demo/inventory_service/startup"
	cfg "github.com/RadisaMilovcevic/bsep_xws/microservices_demo/inventory_service/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
