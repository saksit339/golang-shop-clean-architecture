package entities

import "time"

type Player struct {
	Id        string      `gorm:"primaryKey;type:varchar(64);"`
	Email     string      `gorm:"type:varchar(128);unique;not null;"`
	Name      string      `gorm:"type:varchar(128);not null;"`
	Avatar    string      `gorm:"type:varchar(256);not null;default:'';"`
	Inventory []Inventory `gorm:"foreignKey:PlayerId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time   `gorm:"autoCreateTime;not null;"`
	UpdatedAt time.Time   `gorm:"autoUpdateTime;not null;"`
}
