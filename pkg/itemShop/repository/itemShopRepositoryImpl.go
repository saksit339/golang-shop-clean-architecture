package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/saksit339/isekai-shop-api-tutorial/entities"
	_itemShopExecptions "github.com/saksit339/isekai-shop-api-tutorial/pkg/itemShop/exception"
	_itemShopModel "github.com/saksit339/isekai-shop-api-tutorial/pkg/itemShop/model"

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

func (r *itemShopRepositoryImpl) Listing(itemfilter *_itemShopModel.ItemFilter) ([]*entities.Item, error) {
	items := make([]*entities.Item, 0)

	query := r.db.Model(&entities.Item{})
	if itemfilter.Name != "" {
		query = query.Where("name LIKE ?", "%"+itemfilter.Name+"%")
	}
	if itemfilter.Description != "" {
		query = query.Where("description LIKE ?", "%"+itemfilter.Description+"%")
	}
	query = query.Where("is_archived = ?", false)

	offset := (itemfilter.Pagination.Page - 1) * itemfilter.Pagination.Size

	if err := query.Offset(int(offset)).Limit(int(itemfilter.Size)).Find(&items).Error; err != nil {
		r.logger.Print("Error fetching items")
		return nil, &_itemShopExecptions.ItemListing{}
	}
	return items, nil
}
