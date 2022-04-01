package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/xvbnm48/go-medium-ecomerce/model"
)

type UserRepository interface {
	AddUser(model.User) (model.User, error)
	GetUser(int) (model.User, error)
	GetByEmail(string) (model.User, error)
	GetAllUser() ([]model.User, error)
	UpdateUser(model.User) (model.User, error)
	DeleteUser(model.User) (model.User, error)
	GetProductOrdered(int) ([]model.Order, error)
}

type userRepository struct {
	connection *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{
		connection: DB(),
	}
}

func (db *userRepository) GetUser(id int) (user model.User, err error) {
	return user, db.connection.First(&user, id).Error
}

func (db *userRepository) GetByEmail(email string) (user model.User, err error) {
	return user, db.connection.Where("email = ?", email).First(&user).Error
}

func (db *userRepository) GetAllUser() (users []model.User, err error) {
	return users, db.connection.Find(&users).Error
}

func (db *userRepository) AddUser(user model.User) (model.User, error) {
	return user, db.connection.Create(&user).Error
}

func (db *userRepository) UpdateUser(user model.User) (model.User, error) {
	if err := db.connection.First(&user, user.ID).Error; err != nil {
		return user, err
	}

	return user, db.connection.Model(&user).Updates(&user).Error
}

func (db *userRepository) DeleteUser(user model.User) (model.User, error) {
	if err := db.connection.First(&user, user.ID).Error; err != nil {
		return user, err
	}

	return user, db.connection.Delete(&user).Error
}

func (db *userRepository) GetProductOrdered(userID int) (orders []model.Order, err error) {
	return orders, db.connection.Where("user_ID = ?", userID).First(&orders).Error
}
