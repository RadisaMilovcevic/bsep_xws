package services

import (
	account "github.com/RadisaMilovcevic/bsep_xws/microservices_demo/common/proto/account_service"
	catalogue "github.com/RadisaMilovcevic/bsep_xws/microservices_demo/common/proto/catalogue_service"
	ordering "github.com/RadisaMilovcevic/bsep_xws/microservices_demo/common/proto/ordering_service"
	shipping "github.com/RadisaMilovcevic/bsep_xws/microservices_demo/common/proto/shipping_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewCatalogueClient(address string) catalogue.CatalogueServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Catalogue service: %v", err)
	}
	return catalogue.NewCatalogueServiceClient(conn)
}

func NewOrderingClient(address string) ordering.OrderingServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Ordering service: %v", err)
	}
	return ordering.NewOrderingServiceClient(conn)
}

func NewShippingClient(address string) shipping.ShippingServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Shipping service: %v", err)
	}
	return shipping.NewShippingServiceClient(conn)
}

func NewAccountClient(address string) account.AccountServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Shipping service: %v", err)
	}
	return account.NewAccountServiceClient(conn)
}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
