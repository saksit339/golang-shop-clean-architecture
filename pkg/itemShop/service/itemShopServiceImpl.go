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

func (s *itemShopServiceImpl) Listing(itemfilter *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult, error) {
	itemList, err := s.itemShopRepository.Listing(itemfilter)
	if err != nil {
		return nil, err
	}

	itemCounting, err := s.itemShopRepository.Counting(itemfilter)
	if err != nil {
		return nil, err
	}

	totalPage := s.totalPageCalculation(itemCounting, itemfilter.Size)

	itemModelList := make([]_itemShopModel.Item, 0)
	for _, item := range itemList {
		itemModel := item.ToItemModel()
		itemModelList = append(itemModelList, *itemModel)
	}

	result := _itemShopModel.ItemResult{
		Items: itemModelList,
	}

	return &result, nil
}

func (itemShopServiceImpl) totalPageCalculation(totalIten int64, size int64) int64 {
	total := totalIten / size
	if totalIten%size != 0 {
		total++
	}
	return total
}
