package user

import (
	"go.mongodb.org/mongo-driver/mongo"
	"mikaellemos.com.br/dload/src/config"
	"sync"
)

const COLLECTION_NAME = "users"

var once *sync.Once
var collection *mongo.Collection

func Collection() *mongo.Collection {
	db := config.Client()
	collection = db.Collection(COLLECTION_NAME)

	return collection
}

//func StartCollection() {
//	//once.Do(func() {
//	//	db := config.Client()
//	//	collection = db.Collection(COLLECTION_NAME)
//	//})
//}
