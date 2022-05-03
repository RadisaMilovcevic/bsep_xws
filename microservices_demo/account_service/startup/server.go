package startup

import (
	"account_service/application"
	"account_service/domain"
	"account_service/infrastructure/api"
	"account_service/infrastructure/persistence"
	"account_service/startup/config"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
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

	productService := server.initAccountService(accountStore)

	productHandler := server.initAccountHandler(productService)

	server.startGrpcServer(productHandler)
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
