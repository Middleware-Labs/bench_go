package main

import (
	"log"

	"github.com/gin-gonic/gin"

	config "bench_go/config"
	routes "bench_go/routes"
)

func main() {
	// Connect DB
	config.Connect()

	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	routes.Routes(router)

	log.Fatal(router.Run(":4747"))
}
