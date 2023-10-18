package migrate

import (
	"mikaellemos.com.br/dload/src/config"
	"mikaellemos.com.br/dload/src/model"
)

func Migrate() {
	config.PostClient().AutoMigrate(&model.User{})
}
