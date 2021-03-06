package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	Id       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
}
