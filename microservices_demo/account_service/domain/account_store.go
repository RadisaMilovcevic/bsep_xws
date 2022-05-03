package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type AccountStore interface {
	Get(id primitive.ObjectID) (*Account, error)
	GetAll() ([]*Account, error)
	Insert(account *Account) error
	DeleteAll()
}
