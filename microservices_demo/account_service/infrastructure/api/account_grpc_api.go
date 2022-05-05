package api

import (
	"context"
	"github.com/RadisaMilovcevic/bsep_xws/microservices_demo/account_service/application"
	pb "github.com/RadisaMilovcevic/bsep_xws/microservices_demo/common/proto/account_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountHandler struct {
	pb.UnimplementedAccountServiceServer
	service *application.AccountService
}

func NewAccountHandler(service *application.AccountService) *AccountHandler {
	return &AccountHandler{
		service: service,
	}
}

func (handler *AccountHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	account, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	accountPb := mapAccount(account)
	response := &pb.GetResponse{
		Account: accountPb,
	}
	return response, nil
}

func (handler *AccountHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	accounts, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Accounts: []*pb.Account{},
	}
	for _, account := range accounts {
		current := mapAccount(account)
		response.Accounts = append(response.Accounts, current)
	}
	return response, nil
}

func (handler *AccountHandler) CreateAccount(ctx context.Context, request *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	account := request.Account
	err := handler.service.Create(mapNewAccount(account))
	if err != nil {
		return nil, err
	}
	return &pb.CreateAccountResponse{
		Account: account,
	}, nil
}
