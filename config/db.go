package config

import (
	"log"
	"os"
	"github.com/go-pg/pg/v9"

	controllers "github.com/cavdy-play/go_db/controllers"
)

// Connecting to db
func Connect() *pg.DB {
	opts := &pg.Options{
		User: "postgres",
		Password: "bcv8TBzS3vqbGciO2J73gwLYaoYbGciO2J73gwLopaopa4eCAue3vqonAue3vqo",
		Addr: "localhost:5432",
		Database: "benchdb",
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	controllers.CreateTodoTable(db)
	controllers.InitiateDB(db)
	return db
}
