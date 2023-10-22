package repository

import (
	"math"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"mikaellemos.com.br/dload/src/config"
	"mikaellemos.com.br/dload/src/model"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository() *Repository {
	db := config.PostClient()
	return &Repository{db}
}

func NewRepositoryDB(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) Insert(u *model.User) (*model.User, error) {
	result := r.db.Create(&u)
	if result.Error != nil {
		return nil, result.Error
	}
	return u, nil
}

func (r *Repository) Update(u *model.User) (*model.User, error) {
	result := r.db.Model(&u).Updates(&model.User{
		Name:     u.Name,
		Birthday: u.Birthday,
	})
	if result.Error != nil {
		return nil, result.Error
	}

	return u, nil
}

func (r *Repository) Delete(id uuid.UUID) error {
	p := model.User{ID: id}
	result := r.db.Delete(&p)
	return result.Error
}

func (r *Repository) FindById(id uuid.UUID) (*model.User, error) {
	var user model.User
	result := r.db.First(&user, id)

	if result.Error != nil {
		logrus.Info(result.Error)
		return nil, result.Error
	}

	return &user, nil
}

func (r *Repository) Paginate(user *model.User, page int, limit int) *model.Paginate {
	var totalRows int64
	var users []model.User
	r.db.Model(&user).Count(&totalRows)

	totalPages := int64(math.Ceil(float64(totalRows) / float64(limit)))
	offset := limit * page

	r.db.Limit(limit).Offset(offset).Find(&users)

	return &model.Paginate{
		NumOfPage: totalPages,
		Page:      page,
		Content:   &users,
		Total:     totalRows,
	}
}

func (r *Repository) FindByEmail(user *model.User, email string) bool {
	r.db.Where("email", email).First(user)
	return user.Email != ""
}
