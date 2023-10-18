package userserv

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"mikaellemos.com.br/dload/src/config"
	"mikaellemos.com.br/dload/src/model"
)

const LIMIT = 10

func List(page int) *[]model.User {
	var users []model.User
	db := config.PostClient()
	offset := page * LIMIT
	db.Limit(LIMIT).Offset(offset).Find(&users)
	return &users
}

func Create(payload []byte) (*model.User, *[]model.ApiError) {
	var _user model.User
	_ = json.Unmarshal(payload, &_user)

	db := config.PostClient()
	db.Create(&_user)

	return &_user, nil
}

func Update(id uint64, payload []byte) *model.User {
	var _user model.User
	_ = json.Unmarshal(payload, &_user)
	_user.ID = id

	db := config.PostClient()
	db.Model(&model.User{}).Where("id", id).Updates(_user)

	return &_user
}

func Remove(id int) {
	db := config.PostClient()
	db.Where("id", id).Delete(&model.User{})

	logrus.Info(fmt.Sprintf("Remove user by id [%d]", id))
}
