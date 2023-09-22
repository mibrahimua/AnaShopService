package service

import (
	"AnaShopService/model"
	"AnaShopService/repository"
)

type ShopService struct {
	productRepository *repository.ProductRepository
}

func NewProductService(productRepository *repository.ProductRepository) *ShopService {
	return &ShopService{productRepository}
}

func (us *ShopService) GetListProductByName(productName string) ([]model.Product, error) {
	return us.productRepository.GetListProductByName(productName)
}
