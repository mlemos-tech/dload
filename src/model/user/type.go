package user

import (
	"context"
	"mikaellemos.com.br/dload/src/config"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//const COLLECTION_NAME = "users"

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name" json:"name" validate:"required"`
	Email     string             `bson:"email" json:"email" validate:"email,required"`
	CreatedAt primitive.DateTime `bson:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at"`
}

func (m User) Create(model interface{}) error {

	db := config.Client()
	collection := db.Collection(COLLECTION_NAME)
	//loc, _ := time.LoadLocation("America/Sao_Paulo")

	m.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	m.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	res, err := collection.InsertOne(context.Background(), model)

	if err != nil {
		logrus.Error(err)
		return err
	}

	m.ID = res.InsertedID.(primitive.ObjectID)
	return nil

}

func (m User) List(model interface{}) []User {

	db := config.Client()
	collection := db.Collection(COLLECTION_NAME)
	ctx := context.Background()
	limit := int64(2)
	skip := int64(2)
	opt := options.FindOptions{Limit: &limit, Skip: &skip}

	list, _ := collection.Find(ctx, bson.D{{}}, &opt)
	var listUser []User

	for list.Next(ctx) {
		var u User
		if err := list.Decode(&u); err != nil {
			continue
		}

		listUser = append(listUser, u)
	}

	return listUser
}
