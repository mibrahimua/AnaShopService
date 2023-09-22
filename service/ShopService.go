package service

import (
	"AnaShopService/repository"
)

type ShopService struct {
	shopRepository *repository.ShopRepository
}

func NewShopService(shopRepository *repository.ShopRepository) *ShopService {
	return &ShopService{shopRepository}
}

func (us *ShopService) PaidCartItems(cartId int) error {
	return us.shopRepository.PaidCartItems(cartId)
}
