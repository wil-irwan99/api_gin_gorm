package usecase

import (
	"api_gin_ref/model"
	"api_gin_ref/repository"
)

type GetProductUseCase interface {
	GetProduct() ([]model.Product, error)
}

type getProductUseCase struct {
	repo repository.ProductRepository
}

func (g *getProductUseCase) GetProduct() ([]model.Product, error) {
	return g.repo.Retrieve()
}

func NewGetProductUseCase(repo repository.ProductRepository) GetProductUseCase {
	return &getProductUseCase{
		repo: repo,
	}
}
