package main

import (
	"fmt"

	"github.com/saksit339/isekai-shop-api-tutorial/config"
	"github.com/saksit339/isekai-shop-api-tutorial/databases"
	"github.com/saksit339/isekai-shop-api-tutorial/entities"
)

func main() {
	conf := config.GetConfig()
	db := databases.NewPostgresDatabase(conf.Database)
	fmt.Println(db.ConnectionGetting())

	playerCoinMigration(db)
	playerMigration(db)
	adminMigration(db)
	itemMigration(db)
	inventoryMigration(db)
	purcheaseHistoryMigration(db)

}
func playerMigration(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.Player{})
}

func adminMigration(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.Admin{})
}

func itemMigration(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.Item{})
}

func inventoryMigration(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.Inventory{})
}

func playerCoinMigration(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.PlayerCoin{})
}

func purcheaseHistoryMigration(db databases.Database) {
	db.ConnectionGetting().Migrator().CreateTable(&entities.PurchaseHistory{})
}
