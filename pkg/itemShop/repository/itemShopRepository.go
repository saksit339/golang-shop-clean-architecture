package repository

import (
	"github.com/saksit339/isekai-shop-api-tutorial/entities"
	_itemShopModel "github.com/saksit339/isekai-shop-api-tutorial/pkg/itemShop/model"
)

type ItemShopRepository interface {
	Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error)
	Counting(itemfilter *_itemShopModel.ItemFilter) (int64, error)
}
