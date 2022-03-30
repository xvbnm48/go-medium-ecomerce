package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/xvbnm48/go-medium-ecomerce/model"
)

type OrderRepository interface {
	OrderProduct(int, int, int) error
}

type orderRepository struct {
	connection *gorm.DB
}

func NewOrderRepository() OrderRepository {
	return &orderRepository{
		connection: DB(),
	}
}

func (db *orderRepository) OrderProduct(userID int, productID int, quantity int) error {
	return db.connection.Create(&model.Order{
		UserID:    uint(userID),
		ProductID: uint(productID),
		Quantity:  quantity,
	}).Error
}
