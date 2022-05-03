package api

import (
	"account_service/application"
	"context"
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
	product, err := handler.service.Get(objectId)
	if err != nil {
		return nil, err
	}
	productPb := mapProduct(product)
	response := &pb.GetResponse{
		Account: accountPb,
	}
	return response, nil
}

func (handler *ProductHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	products, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllResponse{
		Products: []*pb.Product{},
	}
	for _, product := range products {
		current := mapProduct(product)
		response.Products = append(response.Products, current)
	}
	return response, nil
}
