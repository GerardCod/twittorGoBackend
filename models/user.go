package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User is a model representing an user in mongo DB
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name,omitempty"`
	LastName  string             `bson:"last_name" json:"last_name,omitempty"`
	BirthDate time.Time          `bson:"birth_date" json:"birth_date,omitempty"`
	Email     string             `bson:"email" json:"email,omitempty"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar,omitempty"`
	Banner    string             `bson:"banner" json:"banner,omitempty"`
	Biografy  string             `bson:"biografy,omitempty" json:"biografy,omitempty"`
	Location  string             `bson:"location" json:"location,omitempty"`
	WebSite   string             `json:"web_site" json:"web_site,omitempty"`
}
