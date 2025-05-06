package service

import (
	_itemShopRepository "github.com/saksit339/isekai-shop-api-tutorial/pkg/itemShop/repository"

	_itemShopModel "github.com/saksit339/isekai-shop-api-tutorial/pkg/itemShop/model"
)

type itemShopServiceImpl struct {
	itemShopRepository _itemShopRepository.ItemShopRepository
}

func NewItemShopServiceImpl(itemShopService _itemShopRepository.ItemShopRepository) ItemShopService {
	return &itemShopServiceImpl{itemShopService}
}

func (s *itemShopServiceImpl) Listing() ([]*_itemShopModel.Item, error) {
	itemList, err := s.itemShopRepository.Listing()
	if err != nil {
		return nil, err
	}

	itemModelList := make([]*_itemShopModel.Item, 0)
	for _, item := range itemList {
		itemModel := item.ToItemModel()
		itemModelList = append(itemModelList, itemModel)
	}

	return itemModelList, nil
}
