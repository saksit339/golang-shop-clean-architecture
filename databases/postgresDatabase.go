package databases

import (
	"fmt"
	"log"
	"sync"

	"github.com/saksit339/isekai-shop-api-tutorial/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDatabase struct {
	*gorm.DB
}

var postgresDatabaseInstance *PostgresDatabase
var once sync.Once

func NewPostgresDatabase(conf *config.Database) Database {
	once.Do(func() {
		// Updated connection string without 'schema' and replaced with 'search_path'
		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s search_path=%s",
			conf.Host, conf.Port, conf.User, conf.Password, conf.DbName, conf.Sslmode, conf.Schema)
		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database: " + err.Error())
		}

		log.Println("Postgres database connected: " + conf.DbName)

		postgresDatabaseInstance = &PostgresDatabase{DB: conn}
	})
	return postgresDatabaseInstance
}

func (p *PostgresDatabase) ConnectionGetting() *gorm.DB {
	return p.DB
}
