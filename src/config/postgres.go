package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var postgreOnce = sync.Once{}
var postgreInstance *gorm.DB

func ConnectPostgres(db DB) {

	postgreOnce.Do(func() {
		postgreInstance, _ = gorm.Open(postgres.Open(db.Uri), &gorm.Config{})
	})

}

func PostClient() *gorm.DB {
	return postgreInstance
}
