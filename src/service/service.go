package service

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"mikaellemos.com.br/dload/src/model"
	"mikaellemos.com.br/dload/src/repository"
	"time"
)

const LIMIT = 10

func List(page int) *model.Paginate {
	repo := repository.NewRepository()
	return repo.Paginate(&model.User{}, page, LIMIT)
}

func Create(payload []byte) (*model.User, *model.ApiError) {
	var user model.User
	_ = json.Unmarshal(payload, &user)
	repo := repository.NewRepository()
	exists := repo.FindByEmail(&model.User{}, user.Email)

	if exists {
		return nil, &model.ApiError{Error: "Email already exists", Field: "Email"}
	}

	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := repo.Insert(&user)

	if err != nil {
		logrus.Info(err)
	}

	return &user, nil
}

func Update(id uuid.UUID, payload []byte) *model.User {
	var user model.User

	_ = json.Unmarshal(payload, &user)
	repo := repository.NewRepository()
	user.ID = id

	repo.Update(&user)
	return &user
}

func Remove(id uuid.UUID) {
	repo := repository.NewRepository()
	_ = repo.Delete(id)

	logrus.Info(fmt.Sprintf("Remove user by id [%d]", id))
}

func FindById(id uuid.UUID) (*model.User, error) {
	repo := repository.NewRepository()
	return repo.FindById(id)
}
