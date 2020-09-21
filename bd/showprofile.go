package bd

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/GerardCod/twittorGoBackend/models"
)

// ShowProfile retrieves the user's information, given an id.
func ShowProfile(id string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoConnection.Database(DBName)
	collection := db.Collection(UserCollection)

	var profile models.User
	objID, _ := primitive.ObjectIDFromHex(id)
	condition := bson.M{
		"_id": objID,
	}

	err := collection.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""

	if err != nil {
		fmt.Println("Profile not found." + err.Error())
		return profile, err
	}
	return profile, nil
}
