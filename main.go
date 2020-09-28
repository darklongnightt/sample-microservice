package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/darklongnightt/microservice/config"

	"github.com/darklongnightt/microservice/productpage"

	"github.com/darklongnightt/microservice/db"
	"github.com/darklongnightt/microservice/homepage"
	"github.com/darklongnightt/microservice/server"
)

const (
	port = ":8080"
)

func main() {
	// Setup logger
	logger := log.New(os.Stdout, "microservice: ", log.LstdFlags|log.Lshortfile)
	fmt.Printf("Server started on localhost%v\n", port)

	// Get app configs
	config, err := config.GetAppConfig()
	if err != nil {
		logger.Fatalf("Server failed to start: %v", err)
	}

	// Init the db
	db, err := db.Init(config)
	if err != nil {
		logger.Fatalf("Server failed to start: %v", err)
	}
	defer db.Close()

	// Create mux server routings
	mux := http.NewServeMux()

	// Create home handlers with necessary dependencies and setup routes
	homeHandlers := homepage.NewHandlers(logger)
	homeHandlers.SetupRoutes(mux)

	productHandlers := productpage.NewHandlers(logger, db)
	productHandlers.SetupRoutes(mux)

	// Create a new customer server with TLS security
	srv := server.New(mux, port)
	cert := os.Getenv("LOCALHOST_CERT")
	privateKey := os.Getenv("LOCALHOST_PRIVATE_KEY")

	if err := srv.ListenAndServeTLS(cert, privateKey); err != nil {
		logger.Fatalf("Server failed to start: %v", err)
	}
}
