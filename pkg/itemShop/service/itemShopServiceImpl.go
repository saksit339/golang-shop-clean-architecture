package service

import (
	_itemShopRepository "github.com/saksit339/isekai-shop-api-tutorial/pkg/itemShop/repository"
)

type itemShopServiceImpl struct {
	itemShopRepository _itemShopRepository.ItemShopRepository
}

func NewItemShopServiceImpl(itemShopService _itemShopRepository.ItemShopRepository) ItemShopService {
	return &itemShopServiceImpl{itemShopService}
}
