package models

import (
	"time"
	"github.com/google/uuid"
)

type User struct {
	Id uuid.UUID `json:"uer_id"`
	Name string `json:"name"`
	Email string `gorm:"unique" json:"email"`
	Password string `json:"-"`
	Address string `json:"address"`
	Phone int `json:"phone"`
	IsAdmin bool
	CreateAt time.Time `json:"created_at"`
}