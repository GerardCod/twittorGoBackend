package bd

import (
	"context"
	"time"

	"github.com/GerardCod/twittorGoBackend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SignUp create the user in DB
func SignUp(user models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database(DBName)
	collection := db.Collection(UserCollection)

	user.Password, _ = EncryptPassword(user.Password)
	result, err := collection.InsertOne(ctx, user)

	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
