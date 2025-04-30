package main

import (
	"fmt"

	"github.com/saksit339/isekai-shop-api-tutorial/config"
	"github.com/saksit339/isekai-shop-api-tutorial/databases"
	"github.com/saksit339/isekai-shop-api-tutorial/entities"
	"gorm.io/gorm"
)

func main() {
	conf := config.GetConfig()
	db := databases.NewPostgresDatabase(conf.Database)
	fmt.Println(db.ConnectionGetting())

	tx := db.ConnectionGetting().Begin()

	playerCoinMigration(tx)
	playerMigration(tx)
	adminMigration(tx)
	itemMigration(tx)
	inventoryMigration(tx)
	purcheaseHistoryMigration(tx)

	tx.Commit()
	if tx.Error != nil {
		tx.Rollback()
		panic(tx.Error)
	}
}
func playerMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Player{})
}

func adminMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Admin{})
}

func itemMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Item{})
}

func inventoryMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Inventory{})
}

func playerCoinMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.PlayerCoin{})
}

func purcheaseHistoryMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.PurchaseHistory{})
}
