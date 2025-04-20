package entities

import "time"

type Inventory struct {
	Id        uint64    `gorm:"primaryKey;autoIncrement;"`
	PlayerId  string    `gorm:"type:varchar(64);not null;"`
	ItemId    uint64    `gorm:"not null;type:bigint;"`
	IsDeleted bool      `gorm:"default:false;not null;"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null;"`
}
