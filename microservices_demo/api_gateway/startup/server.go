package startup

import (
	"context"
	"fmt"
	"github.com/RadisaMilovcevic/bsep_xws/microservices_demo/api_gateway/infrastructure/api"
	cfg "github.com/RadisaMilovcevic/bsep_xws/microservices_demo/api_gateway/startup/config"
	accountGw "github.com/RadisaMilovcevic/bsep_xws/microservices_demo/common/proto/account_service"
	catalogueGw "github.com/RadisaMilovcevic/bsep_xws/microservices_demo/common/proto/catalogue_service"
	inventoryGw "github.com/RadisaMilovcevic/bsep_xws/microservices_demo/common/proto/inventory_service"
	orderingGw "github.com/RadisaMilovcevic/bsep_xws/microservices_demo/common/proto/ordering_service"
	shippingGw "github.com/RadisaMilovcevic/bsep_xws/microservices_demo/common/proto/shipping_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

type Server struct {
	config *cfg.Config
	mux    *runtime.ServeMux
}

func NewServer(config *cfg.Config) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}
	server.initHandlers()
	server.initCustomHandlers()
	return server
}

func (server *Server) initHandlers() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	catalogueEmdpoint := fmt.Sprintf("%s:%s", server.config.CatalogueHost, server.config.CataloguePort)
	err := catalogueGw.RegisterCatalogueServiceHandlerFromEndpoint(context.TODO(), server.mux, catalogueEmdpoint, opts)
	if err != nil {
		panic(err)
	}
	orderingEmdpoint := fmt.Sprintf("%s:%s", server.config.OrderingHost, server.config.OrderingPort)
	err = orderingGw.RegisterOrderingServiceHandlerFromEndpoint(context.TODO(), server.mux, orderingEmdpoint, opts)
	if err != nil {
		panic(err)
	}
	shippingEmdpoint := fmt.Sprintf("%s:%s", server.config.ShippingHost, server.config.ShippingPort)
	err = shippingGw.RegisterShippingServiceHandlerFromEndpoint(context.TODO(), server.mux, shippingEmdpoint, opts)
	if err != nil {
		panic(err)
	}
	inventoryEmdpoint := fmt.Sprintf("%s:%s", server.config.InventoryHost, server.config.InventoryPort)
	err = inventoryGw.RegisterInventoryServiceHandlerFromEndpoint(context.TODO(), server.mux, inventoryEmdpoint, opts)
	if err != nil {
		panic(err)
	}
	accountEmdpoint := fmt.Sprintf("%s:%s", server.config.AccountHost, server.config.AccountPort)
	err = accountGw.RegisterAccountServiceHandlerFromEndpoint(context.TODO(), server.mux, accountEmdpoint, opts)
	if err != nil {
		panic(err)
	}
}

func (server *Server) initCustomHandlers() {
	catalogueEmdpoint := fmt.Sprintf("%s:%s", server.config.CatalogueHost, server.config.CataloguePort)
	orderingEmdpoint := fmt.Sprintf("%s:%s", server.config.OrderingHost, server.config.OrderingPort)
	shippingEmdpoint := fmt.Sprintf("%s:%s", server.config.ShippingHost, server.config.ShippingPort)
	accountEmdpoint := fmt.Sprintf("%s:%s", server.config.AccountHost, server.config.AccountPort)
	orderingHandler := api.NewOrderingHandler(orderingEmdpoint, catalogueEmdpoint, shippingEmdpoint)
	orderingHandler.Init(server.mux)
	accountHandler := api.NewAccountHandler(accountEmdpoint)
	accountHandler.Init(server.mux)
}

func (server *Server) Start() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), server.mux))
}
