package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required, email" gorm:"unique"`
	Password string `json:"password" binding:"required"`
}

func (User) TableName() string {
	return "users"

}
