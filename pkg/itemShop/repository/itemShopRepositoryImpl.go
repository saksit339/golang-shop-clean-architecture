package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/saksit339/isekai-shop-api-tutorial/entities"
	_itemShopExecptions "github.com/saksit339/isekai-shop-api-tutorial/pkg/itemShop/exception"
	"gorm.io/gorm"
)

type itemShopRepositoryImpl struct {
	db     *gorm.DB
	logger echo.Logger
}

func NewItemShopRepositoryImpl(db *gorm.DB, logger *echo.Logger) ItemShopRepository {
	return &itemShopRepositoryImpl{
		db:     db,
		logger: *logger,
	}
}

func (r *itemShopRepositoryImpl) Listing() ([]*entities.Item, error) {
	items := make([]*entities.Item, 0)
	if err := r.db.Find(&items).Error; err != nil {
		r.logger.Print("Error fetching items")
		return nil, &_itemShopExecptions.ItemListing{}
	}
	return items, nil
}
