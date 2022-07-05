package server

import (
	"log"

	"os"

	"github.com/gin-gonic/gin"
)

func Run() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	// Starts a new Gin instance with no middle-ware
	r := gin.Default()

	// Define handlers
	route(r)

	// Listen and serve on defined port
	log.Printf("Listening on port %s", port)
	r.Run(":" + port)
}
