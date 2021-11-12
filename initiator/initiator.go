package initiator

import (
	"fmt"
	"image/platform/db"
	"log"
	"os"

	ginrouter "image/platform/gin"
)

func DbInitiator() db.DbPlatform {
	// Initaite cockroch platform
	cockroachPlatform := db.Initialize(dbURL)
	// Migrate necessary tables
	err := cockroachPlatform.Migrate()
	if err != nil {
		log.Fatal(err)
	}
	return cockroachPlatform
}

func Initiator() {
	// Initiate db
	dbPlatform := DbInitiator()

	// Initiate image module and
	imageRouter := Image(dbPlatform)

	// Get self host port
	hostPort := os.Getenv("SELF_PORT")
	// hostAddress := os.Getenv("SELF_ADDRESS")
	hostURL := fmt.Sprintf(":" + hostPort)
	//
	domain := os.Getenv("Domain")
	// Get the allowed request origins for the http server
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	// Initiate the http server
	server := ginrouter.Initialize(hostURL, allowedOrigins, imageRouter, domain)

	// Get the handlers from

	// Start the http server
	server.Serve()
}
