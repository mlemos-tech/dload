package model

import (
	"time"
)

type User struct {
	ID        uint64     `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name" validate:"required"`
	Birthday  *time.Time `json:"birthday" validate:"required"`
	Email     *string    `json:"email" validate:"email,required" gorm:"index,unique"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
