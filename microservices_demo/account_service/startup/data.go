package startup

import (
	"github.com/RadisaMilovcevic/bsep_xws/microservices_demo/account_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var accounts = []*domain.Account{
	{
		Id:       getObjectId("623b0cc3a34d25d8567f9f82"),
		Username: "marko",
		Password: "marko",
	},
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
