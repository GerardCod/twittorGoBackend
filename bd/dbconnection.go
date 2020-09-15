package bd

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const password = "6vQicR67scC9LXY"
const dbName = "twittor"

var connectionString = fmt.Sprintf("mongodb+srv://admin:%s@cluster0.vdjb1.mongodb.net/%s?retryWrites=true&w=majority", password, dbName)

//MongoConnection is the connection object for the database
var MongoConnection = DBConnection()

var clientOptions = options.Client().ApplyURI(connectionString)

/*DBConnection does the connection with the database mongo*/
func DBConnection() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Successful connection :)")
	return client
}

/*CheckConnection verifies the connection state*/
func CheckConnection() int {
	err := MongoConnection.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
