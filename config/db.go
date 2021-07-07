package config

import (
	"fmt"
	"log"
	"os"

	//"fmt"
	"github.com/go-pg/pg/v9"

	controllers "bench_go/controllers"
)

// Connecting to db
func Connect(GetEnv map[string]string) *pg.DB {
	fmt.Printf("PG USER env : %s = %s \n", "DB USER", GetEnv["PG_ADDR"])
	fmt.Printf("PG PASS env : %s = %s \n", "DB PASS", GetEnv["PG_DB"])
	opts := &pg.Options{
		User:     GetEnv["PG_USER"],
		Password: GetEnv["PG_PASS"],
		Addr:     GetEnv["PG_ADDR"],
		Database: GetEnv["PG_DB"],
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
