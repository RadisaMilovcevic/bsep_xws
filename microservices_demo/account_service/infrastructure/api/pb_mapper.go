package api

import (
	"account_service/domain"
	pb "github.com/RadisaMilovcevic/bsep_xws/microservices_demo/common/proto/account_service"
)

func mapAccount(account *domain.Account) *pb.Account {
	accountPb := &pb.Account{
		Id:       account.Id.Hex(),
		Username: account.Username,
		Password: account.Password,
	}
	return accountPb
}
