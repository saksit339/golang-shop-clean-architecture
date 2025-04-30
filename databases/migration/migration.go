package main

import (
	"fmt"

	"github.com/saksit339/isekai-shop-api-tutorial/config"
	"github.com/saksit339/isekai-shop-api-tutorial/databases"
)

func main() {
	conf := config.GetConfig()
	db := databases.NewPostgresDatabase(conf.Database)
	fmt.Println(db.ConnectionGetting())
}
