package entities

import "time"

type Item struct {
	Id          uint64    `gorm:"primaryKey;autoIncrement;"`
	AdminId     *string   `gorm:"type:varchar(64);"`
	Name        string    `gorm:"type:varchar(64);unique;not null;"`
	Description string    `gorm:"type:varchar(128);not null;"`
	Pocture     string    `gorm:"type:varchar(256);not null;"`
	Price       uint      `gorm:"not null;"`
	IsArchived  bool      `gorm:"default:false;not null;"`
	CreatedAt   time.Time `gorm:"autoCreateTime;not null;"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime;not null;"`
}
