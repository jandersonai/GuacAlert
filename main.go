//GuacAlert is an Apache Guacamole alerting system that sends messages to a Google Workspace Chat room when a user takes control of a machine and when they release control.
// Users can also check the status of a machine by sending a message to the chat room. If a machine is in use, the system will respond with the user who is currently controlling the machine.
// Users can also enter themselves into a queue to be notified when a machine is available.

package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

func main() {
	// Check for required environment variables
	if dbIP == "" || dbPort == "" || dbUser == "" || dbPass == "" || dbName == "" || chatURL == "" {
		log.Fatal("Missing required environment variables")
	}

	// Check for optional environment variables
	if listenPort == "" {
		listenPort = "8080"
	}

	// Create a new router
	router := NewRouter()

	// Create a new server
	server := &http.Server{
		Addr:         ":" + listenPort,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start the server
	log.Println("Starting server on port " + listenPort)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	// Connect to the database
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

}
