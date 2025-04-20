package entities

import "time"

type PlayerCoin struct {
	Id        uint64    `gorm:"primaryKey;autoIncrement;"`
	PlayerId  string    `gorm:"type:varchar(64);not null;"`
	Amount    int64     `gorm:"not null;"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null;"`
}
