package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	config "bench_go/config"
	routes "bench_go/routes"
)

var PG_USER string
var PG_PASS string

func main() {

	// load .env file from given path
	// we keep it empty it will load .env from current directory
	var GetEnv map[string]string
	GetEnv, err := godotenv.Read(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// getting env variables PG_USER and DB_HOST
	//fmt.Printf("godotenv : %s = %s \n", "DB USER", GetEnv["PG_USER"])
	//fmt.Printf("godotenv : %s = %s \n", "DB PASS", GetEnv["PG_PASS"])

	// Connect DB
	config.Connect(GetEnv)

	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	routes.Routes(router)

	log.Fatal(router.Run(":4747"))
}
