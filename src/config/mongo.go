package config

import (
	"context"
	"github.com/sirupsen/logrus"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var once sync.Once
var instance *mongo.Client
var dbInstance *mongo.Database

func Connect(db *DB) {
	once.Do(func() {

		logrus.Info(db.Uri)
		clientOptions := options.Client().ApplyURI(db.Uri)
		instance, err := mongo.Connect(context.Background(), clientOptions)

		if err != nil {
			log.Fatal(err)
		}

		instance.Ping(context.Background(), nil)
		dbInstance = instance.Database(db.Database)
	})
}

func Client() *mongo.Database {
	return dbInstance
}
