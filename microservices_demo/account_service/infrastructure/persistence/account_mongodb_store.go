package persistence

import (
	"account_service/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "account"
	COLLECTION = "account"
)

type AccountMongoDBStore struct {
	accounts *mongo.Collection
}

func NewAccountMongoDBStore(client *mongo.Client) domain.AccountStore {
	accounts := client.Database(DATABASE).Collection(COLLECTION)
	return &AccountMongoDBStore{
		accounts: accounts,
	}
}

func (store *AccountMongoDBStore) Get(id primitive.ObjectID) (*domain.Account, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *AccountMongoDBStore) GetAll() ([]*domain.Account, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *AccountMongoDBStore) Insert(account *domain.Account) error {
	result, err := store.accounts.InsertOne(context.TODO(), account)
	if err != nil {
		return err
	}
	account.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *AccountMongoDBStore) DeleteAll() {
	store.accounts.DeleteMany(context.TODO(), bson.D{{}})
}

func (store *AccountMongoDBStore) filterOne(filter interface{}) (account *domain.Account, err error) {
	result := store.accounts.FindOne(context.TODO(), filter)
	err = result.Decode(&account)
	return
}

func (store *AccountMongoDBStore) filter(filter interface{}) ([]*domain.Account, error) {
	cursor, err := store.accounts.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func decode(cursor *mongo.Cursor) (accounts []*domain.Account, err error) {
	for cursor.Next(context.TODO()) {
		var account domain.Account
		err = cursor.Decode(&account)
		if err != nil {
			return
		}
		accounts = append(accounts, &account)
	}
	err = cursor.Err()
	return
}
