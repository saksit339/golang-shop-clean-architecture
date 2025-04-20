package entities

type PurchaseHistory struct {
	Id              uint64 `gorm:"primaryKey;autoIncrement;"`
	PlayerId        string `gorm:"type:varchar(64);not null;"`
	ItemId          uint64 `gorm:"not null;type:bigint;"`
	ItemName        string `gorm:"type:varchar(64);not null;"`
	ItemDescription string `gorm:"type:varchar(128);not null;"`
	ItemPrice       uint   `gorm:"not null;"`
	ItemPicture     string `gorm:"type:varchar(128);not null;"`
	Quantity        uint   `gorm:"not null;"`
	CreatedAt       string `gorm:"autoCreateTime;not null;"`
}
