package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName string `gorm:"not null;type:varchar(100);index:unique"`
	Email    string `gorm:"not null;type:varchar(100);index:unique"`
	Password string `gorm:"not null;type:varchar(100)"`
	Role     uint   `gorm:"not null;"`
}

type Category struct {
	gorm.Model
	Type  string
	Tasks []Task
}

type Task struct {
	gorm.Model
	Title       string
	Description string
	Status      bool
	UserID      uint
	User        User

	CategoryID uint
	Category   Category
}
