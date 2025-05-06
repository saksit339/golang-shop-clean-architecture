package main

import (
	"github.com/saksit339/isekai-shop-api-tutorial/config"
	"github.com/saksit339/isekai-shop-api-tutorial/databases"
	"github.com/saksit339/isekai-shop-api-tutorial/server"
)

func main() {
	conf := config.GetConfig()
	db := databases.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db.ConnectionGetting())
	server.Start()
}
