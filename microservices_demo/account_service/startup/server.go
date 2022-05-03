package startup

import (
	"fmt"
	"github.com/RadisaMilovcevic/bsep_xws/microservices_demo/account_service/application"
	"github.com/RadisaMilovcevic/bsep_xws/microservices_demo/account_service/domain"
	"github.com/RadisaMilovcevic/bsep_xws/microservices_demo/account_service/infrastructure/api"
	"github.com/RadisaMilovcevic/bsep_xws/microservices_demo/account_service/infrastructure/persistence"
	"github.com/RadisaMilovcevic/bsep_xws/microservices_demo/account_service/startup/config"
	account "github.com/RadisaMilovcevic/bsep_xws/microservices_demo/common/proto/account_service"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	accountStore := server.initAccountStore(mongoClient)

	accountService := server.initAccountService(accountStore)

	accountHandler := server.initAccountHandler(accountService)

	server.startGrpcServer(accountHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.AccountDBHost, server.config.AccountDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initAccountStore(client *mongo.Client) domain.AccountStore {
	store := persistence.NewAccountMongoDBStore(client)
	store.DeleteAll()
	for _, account := range accounts {
		err := store.Insert(account)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initAccountService(store domain.AccountStore) *application.AccountService {
	return application.NewAccountService(store)
}

func (server *Server) initAccountHandler(service *application.AccountService) *api.AccountHandler {
	return api.NewAccountHandler(service)
}

func (server *Server) startGrpcServer(accountHandler *api.AccountHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	account.RegisterAccountServiceServer(grpcServer, accountHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
