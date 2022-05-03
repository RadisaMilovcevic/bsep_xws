package application

import (
	"github.com/RadisaMilovcevic/bsep_xws/microservices_demo/account_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountService struct {
	store domain.AccountStore
}

func NewAccountService(store domain.AccountStore) *AccountService {
	return &AccountService{
		store: store,
	}
}

func (service *AccountService) Get(id primitive.ObjectID) (*domain.Account, error) {
	return service.store.Get(id)
}

func (service *AccountService) GetAll() ([]*domain.Account, error) {
	return service.store.GetAll()
}

func (service *AccountService) Create(account *domain.Account, address string) error {
	err := service.store.Insert(account)
	if err != nil {
		return err
	}
	if err != nil {
		_ = service.store.UpdateStatus(account)
		return err
	}
	return nil
}
