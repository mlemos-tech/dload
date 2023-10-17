package model

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name" json:"name" validate:"required"`
	Email     string             `bson:"email" json:"email" validate:"email,required"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func (m User) Create(ctx context.Context, db *mongo.Database, collectionName string, model interface{}) error {

	collection := db.Collection(collectionName)
	loc, _ := time.LoadLocation("America/Sao_Paulo")

	m.CreatedAt = time.Now().In(loc)
	m.UpdatedAt = time.Now().In(loc)

	logrus.Info("fdnkdnfkndng jnj" + time.Now().In(loc).String())

	res, err := collection.InsertOne(ctx, model)

	if err != nil {
		logrus.Error(err)
		return err
	}

	m.ID = res.InsertedID.(primitive.ObjectID)
	return nil

}

func (m User) List(ctx context.Context, db *mongo.Database, collectionName string, model interface{}) []User {

	collection := db.Collection(collectionName)
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
