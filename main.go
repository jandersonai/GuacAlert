package main

import (
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

func main() {
	// Print the starting message
	println("Starting GuacAlert. Beep Boop.")

	// Check for required environment variables using a switch
	switch {
	case os.Getenv("GUAC_URL") == "":
		log.Fatal("GUAC_URL is required")
	case os.Getenv("GUAC_USER") == "":
		log.Fatal("GUAC_USER is required")
	case os.Getenv("GUAC_PASS") == "":
		log.Fatal("GUAC_PASS is required")
	case os.Getenv("GUAC_DATASOURCE") == "":
		log.Fatal("GUAC_DATASOURCE is required")
	case os.Getenv("CHAT_HOOK") == "":
		log.Fatal("CHAT_HOOK variable is not set")
	}

	// Generate a token
	println("Generating Token...")
	Token = GenerateToken()
	if Token == "" {
		println("Failed to generate token")
	}
	// Truncate the middle of the token for logging purposes
	println("Token Generation Successful: " + Token[:5] + "..." + Token[len(Token)-5:])

	for {
		// Get the connections from the Guacamole API
		connections := GetConnections(Token)
		if connections == nil {
			return
		}

		// Loop through the queue and remove any connections that are no longer active
		for i := 0; i < len(queue); i++ {
			if !QueueContains(connections, queue[i]) {
				SendMessage(queue[i].Username + " has disconnected from " + queue[i].ServerName)
				queue = append(queue[:i], queue[i+1:]...)
			}
		}

		// Loop through the connections and add them to the queue if they are not already in the queue
		for _, connection := range connections {
			server := GetServer(connection.ConnectionIdentifier, Token)
			if len(server.Name) == 0 {
				return
			}

			// Create a ConnectionDetail struct to store the connection details
			connectionDetail := ConnectionDetail{
				ServerName:   server.Name,
				Username:     connection.Username,
				ConnectionID: connection.ConnectionIdentifier,
			}

			// If the connection does not exist in the queue, add it
			if !ConnectionExists(queue, connectionDetail) {
				SendMessage(connection.Username + " is now connected to " + server.Name)
				queue = append(queue, connectionDetail)
				println("Connection added to Active List: Server: " + server.Name + ", User: " + connection.Username)
			}
		}

		// If the queue is not empty, print the contents of the queue
		if len(queue) > 0 {
			println("Active Connections:")
			for _, connection := range queue {
				println("Server: " + connection.ServerName + ", User: " + connection.Username)
			}
		}

		// Sleep for 5 seconds
		time.Sleep(5 * time.Second)
	}
}
