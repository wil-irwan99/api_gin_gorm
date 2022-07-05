package repository

import (
	"api_gin_ref/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Add(newProduct *model.Product) error
	Retrieve() ([]model.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func (p *productRepository) Add(newProduct *model.Product) error {
	err := p.db.Create(&newProduct).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *productRepository) Retrieve() ([]model.Product, error) {
	var products []model.Product
	err := p.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	repo := new(productRepository)
	repo.db = db
	return repo
}
