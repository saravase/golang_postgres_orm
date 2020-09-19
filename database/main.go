package database

import (
	"log"
	"os"
	"time"

	plant "golang_postgresql_model/plant"

	pg "github.com/go-pg/pg"
)

func Connect() *pg.DB {

	// Configure database options
	options := &pg.Options{
		User:         "xxxxx",
		Password:     "xxxxx@2207",
		Addr:         "localhost:5432",
		Database:     "xxxxx",
		DialTimeout:  30 * time.Second,
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
		IdleTimeout:  30 * time.Minute,
		MaxConnAge:   1 * time.Minute,
		PoolSize:     20,
	}

	// Create connnection
	var db *pg.DB = pg.Connect(options)
	if db == nil {
		log.Printf("Failed to connect the database. %v\n", db)
		os.Exit(100)
	}
	log.Printf("Database connected successfully.\n")

	//Create Table
	plant.CreatePlantInfoTable(db)

	/*Close connnection
	closeError := db.Close()
	if closeError != nil {
		log.Printf("Error while close database connection. %v\n", closeError)
		os.Exit(100)
	}

	log.Printf("Database connection closed successfully. \n")*/
	return db
}
