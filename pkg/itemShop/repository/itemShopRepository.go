package repository

import "github.com/saksit339/isekai-shop-api-tutorial/entities"

type ItemShopRepository interface {
	Listing() ([]*entities.Item, error)
}
