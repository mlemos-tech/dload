package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" validate:"required"`
	Birthday  string    `json:"birthday" validate:"required"`
	Email     string    `json:"email" validate:"email,required" gorm:"index,unique"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
