package bd

import (
	"context"
	"time"

	"github.com/GerardCod/twittorGoBackend/models"
	"go.mongodb.org/mongo-driver/bson"
)

//UserAlreadyExists checks if the is an user with the given email
func UserAlreadyExists(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database(DBName)
	collection := db.Collection(UserCollection)

	condition := bson.M{"email": email}

	var result models.User

	err := collection.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}
