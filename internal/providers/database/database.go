package database

import (
	"fmt"
	"github.com/DevtronLabs/CatPicHub/internal/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var (
	client     *gorm.DB
	clientTest *gorm.DB
)

// Initialize will initialize the connection to the dialect
func Initialize() {
	log.Println("Initializing database...")
	dbConfig := config.GetConfig().Database
	fmt.Print(dbConfig)
	client = connect(config.GetConfig().Database)
}

// Client will return the database client
func Client() *gorm.DB {
	return client
}

func connect(database config.Database) *gorm.DB {
	var err error

	db, err := gorm.Open(database.Dialect, database.URL())

	if err != nil {
		log.Println("Error occurred while initialising database ", err)
		panic("database connection failure")
	}

	// This will prevent update or delete without where clause
	db.BlockGlobalUpdate(true)

	return db
}

func TestClient() *gorm.DB {
	return clientTest
}
