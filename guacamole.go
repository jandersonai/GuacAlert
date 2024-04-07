package main

// DatabaseWatch Implement a function to watch the database for changes
func DatabaseWatch() {
	// Start the loop
	for {
		// Pull connection history from the database
		connections := connectDB
		for _, connection := range connections {
			// Check if the connection is active
			if connection.active {
				// Check if the connection has changed
				if connection.changed {
					// Send a message to the chat room
					SendMessage(connection)
					// Reset the connection changed flag
					connection.changed = false
				}
			}
		}
	}
}
