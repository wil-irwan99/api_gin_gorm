package manager

import "api_gin_ref/usecase"

type UseCaseManager interface {
	CreateProductUseCase() usecase.CreateProductUseCase
	GetProductUseCase() usecase.GetProductUseCase
}

type useCaseManager struct {
	repoManager RepositoryManager
}

func (u *useCaseManager) CreateProductUseCase() usecase.CreateProductUseCase {
	return usecase.NewCreateProductUseCase(u.repoManager.ProductRepo())
}

func (u *useCaseManager) GetProductUseCase() usecase.GetProductUseCase {
	return usecase.NewGetProductUseCase(u.repoManager.ProductRepo())
}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}
