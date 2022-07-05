package manager

import (
	"api_gin_ref/repository"
)

type RepositoryManager interface {
	ProductRepo() repository.ProductRepository
}

type repositoryManager struct {
	infra Infra
}

func (r *repositoryManager) ProductRepo() repository.ProductRepository {
	return repository.NewProductRepository(r.infra.SqlDb())
}

func NewRepositoryManager(infra Infra) RepositoryManager {
	return &repositoryManager{
		infra: infra,
	}
}
