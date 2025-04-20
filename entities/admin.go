package entities

import "time"

type Admin struct {
	Id        string    `gorm:"primaryKey;type:varchar(64);"`
	Item      []Item    `gorm:"foreignKey:AdminId;references:Id;contraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Email     string    `gorm:"type:varchar(128);unique;not null;"`
	Name      string    `gorm:"type:varchar(128);not null;"`
	Avatar    string    `gorm:"type:varchar(256);not null;default:'';"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null;"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;not null;"`
}
