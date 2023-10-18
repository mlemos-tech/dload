package userserv

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mikaellemos.com.br/dload/src/model"
	"mikaellemos.com.br/dload/src/model/user"
	"mikaellemos.com.br/dload/src/service/userserv/validate"
	"time"
)

const LIMIT = 10

func List(page int64) *[]user.User {
	coll := user.Collection()
	ctx := context.Background()
	limit := int64(LIMIT)
	offset := page * LIMIT
	opt := options.FindOptions{Limit: &limit, Skip: &offset}

	list, _ := coll.Find(ctx, bson.D{{}}, &opt)
	listUser := []user.User{}

	for list.Next(ctx) {
		var _user user.User
		if err := list.Decode(&_user); err != nil {
			continue
		}

		listUser = append(listUser, _user)
	}

	return &listUser
}

func Create(payload []byte) (*user.User, *[]model.ApiError) {
	_user, err := validate.Validate(payload)

	if err != nil {
		return nil, err
	}

	coll := user.Collection()
	res, _ := coll.InsertOne(context.Background(), _user)
	_user.ID = res.InsertedID.(primitive.ObjectID)

	return _user, nil
}

func Update(id string, payload []byte) (*user.User, *[]model.ApiError) {
	model, err := validate.Validate(payload)
	ctx := context.Background()

	if err != nil {
		return nil, err
	}

	coll := user.Collection()
	mongoid, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": bson.M{"$eq": mongoid}}
	model.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	_, er := coll.UpdateOne(ctx, filter, &model)

	if er != nil {
		logrus.Info(er)
	}

	var u user.User
	coll.FindOne(ctx, bson.D{{"ID": "mongoid"}}).Decode(&u)

	//logrus.Info(res.UpsertedID)
	//model.ID = res.UpsertedID.(primitive.ObjectID)

	return &u, nil
}

func Remove(id string) {
	coll := user.Collection()
	filter := bson.D{{"_id", id}}
	coll.DeleteOne(context.Background(), filter)
}
