package application

import (
	"account_service/domain"
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
