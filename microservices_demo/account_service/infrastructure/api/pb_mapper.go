package api

import (
	"github.com/RadisaMilovcevic/bsep_xws/microservices_demo/account_service/domain"
	pb "github.com/RadisaMilovcevic/bsep_xws/microservices_demo/common/proto/account_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapAccount(account *domain.Account) *pb.Account {
	accountPb := &pb.Account{
		Id:       account.Id.Hex(),
		Username: account.Username,
		Password: account.Password,
	}
	return accountPb
}

func mapNewAccount(accountPb *pb.Account) *domain.Account {
	id, err := primitive.ObjectIDFromHex(accountPb.Id)
	account := &domain.Account{
		Id:       id,
		Username: accountPb.Username,
		Password: accountPb.Password,
	}
	if err != nil {
	}
	return account
}
