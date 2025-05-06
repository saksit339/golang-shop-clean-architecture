package entities

import (
	"time"

	_itemShopModel "github.com/saksit339/isekai-shop-api-tutorial/pkg/itemShop/model"
)

type Item struct {
	Id          uint64    `gorm:"primaryKey;autoIncrement;"`
	AdminId     *string   `gorm:"type:varchar(64);"`
	Name        string    `gorm:"type:varchar(64);unique;not null;"`
	Description string    `gorm:"type:varchar(128);not null;"`
	Picture     string    `gorm:"type:varchar(256);not null;"`
	Price       uint      `gorm:"not null;"`
	IsArchived  bool      `gorm:"default:false;not null;"`
	CreatedAt   time.Time `gorm:"autoCreateTime;not null;"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime;not null;"`
}

func (i *Item) ToItemModel() *_itemShopModel.Item {
	return &_itemShopModel.Item{
		Id:          i.Id,
		Name:        i.Name,
		Description: i.Description,
		Picture:     i.Picture,
		Price:       i.Price,
		IsArchived:  i.IsArchived,
	}
}
