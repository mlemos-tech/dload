package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var postgreOnce = sync.Once{}
var postgreInstance *gorm.DB

func ConnectPostgres(uri string) {

	postgreOnce.Do(func() {
		postgreInstance, _ = gorm.Open(postgres.Open(uri), &gorm.Config{})
	})

}

func PostClient() *gorm.DB {
	return postgreInstance
}
