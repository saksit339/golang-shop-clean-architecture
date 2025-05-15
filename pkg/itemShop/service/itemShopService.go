package service

import (
	_itemShopModel "github.com/saksit339/isekai-shop-api-tutorial/pkg/itemShop/model"
)

type ItemShopService interface {
	Listing(itemFilter *_itemShopModel.ItemFilter) (*_itemShopModel.ItemResult, error)
}
