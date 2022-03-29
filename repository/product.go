package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/xvbnm48/go-medium-ecomerce/model"
)

type ProductRepository interface {
	GetProduct(int) (model.Product, error)
	GetAllProduct() ([]model.Product, error)
	AddProduct(model.Product) (model.Product, error)
	UpdateProduct(model.Product) (model.Product, error)
	DeleteProduct(model.Product) (model.Product, error)
}

type productRepository struct {
	connection *gorm.DB
}

func NewProductRepository() ProductRepository {
	return &productRepository{
		connection: DB(),
	}
}

func (db *productRepository) GetProduct(id int) (product model.Product, err error) {
	return product, db.connection.First(&product, id).Error
}

func (db *productRepository) GetAllProduct() (products []model.Product, err error) {
	return products, db.connection.Find(&products).Error
}

func (db *productRepository) AddProduct(product model.Product) (model.Product, error) {
	return product, db.connection.Create(&product).Error
}

func (db *productRepository) UpdateProduct(product model.Product) (model.Product, error) {
	return product, db.connection.Save(&product).Error
}

func (db *productRepository) DeleteProduct(product model.Product) (model.Product, error) {
	return product, db.connection.Delete(&product).Error
}
