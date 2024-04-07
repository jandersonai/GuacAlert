package main

// Watch all messages in the chat room and respond to messages that match the following commands:
// !want <servername> - Check the status of a server and if it is already in use add the user to the queue for that server
// !status <servername> - Check the status of the server. If the server is in use, respond with the user who is currently controlling the server, otherwise respond with the server is available
// !remove <servername> - Remove the user from the queue for that server if they are in the queue
// !help - Display a list of available commands and their descriptions
// !list - Display a list of servers and their status (in use or available)
// !who <servername> - Display the user who is currently controlling the server

// ParseMessage Implement a function to parse the message and respond to the appropriate command
func ParseMessage(message string) {
	// Implement the function to parse the message and respond to the appropriate command
}

// AddToQueue Implement a function to add a user to the queue for a server
func AddToQueue(server string, user string) {
	// Implement the function to add a user to the queue for a server
}

// RemoveFromQueue Implement a function to remove a user from the queue for a server
func RemoveFromQueue(server string, user string) {
	// Implement the function to remove a user from the queue for a server
}

// CheckStatus Implement a function to check the status of a server
func CheckStatus(server string) {
	// Implement the function to check the status of a server
}

// ListServers Implement a function to display a list of servers and their status
func ListServers() {
	// Implement the function to display a list of servers and their status
}

// WhoIsControlling Implement a function to display the user who is currently controlling a server
func WhoIsControlling(server string) {
	// Implement the function to display the user who is currently controlling a server
}
