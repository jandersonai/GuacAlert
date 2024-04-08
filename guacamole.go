package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// GenerateToken Generate a token for the user to use to authenticate with the Guacamole API
func GenerateToken() string {
	// Query the Guacamole API using username and password to generate a token
	request, err := http.NewRequest("POST", guacURL+"/api/tokens", authBody)
	if err != nil {
		return ""
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Accept", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return ""
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)
	token, err := io.ReadAll(response.Body)
	if err != nil {
		return ""
	}
	// grab the authToken from the json response
	authToken := strings.Split(string(token), "\"")[3]
	return authToken
}

// RefreshToken Refresh the token to authenticate with the Guacamole API
func RefreshToken() string {
	request, err := http.NewRequest("PUT", guacURL+"/api/tokens/"+Token, nil)
	if err != nil {
		return ""
	}
	request.Header.Set("Accept", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return ""
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)
	token, err := io.ReadAll(response.Body)
	if err != nil {
		return ""
	}
	// grab the authToken from the json response
	authToken := strings.Split(string(token), "\"")[3]
	Token = authToken
	return authToken
}

// GetConnections Pull the list of active connections from the Guacamole API
func GetConnections(token string) map[string]Connection {
	request, err := http.NewRequest("GET", guacURL+"/api/session/data/postgresql/activeConnections?token="+token, nil)
	if err != nil {
		println("Error creating request")
		return nil
	}
	request.Header.Set("Accept", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		println("Error sending request")
		return nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			println("Error closing body")
			return
		}
	}(response.Body)

	// Check if the token has expired
	if response.StatusCode == http.StatusUnauthorized {
		// Refresh the token
		token = RefreshToken()
		if token == "" {
			return nil
		}

		response, err = client.Do(request)
		if err != nil {
			return nil
		}
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		println("Error reading body")
		return nil
	}

	var rawConnections map[string]json.RawMessage
	err = json.Unmarshal(body, &rawConnections)
	if err != nil {
		println("Error unmarshalling json")
		return nil
	}

	connections := make(map[string]Connection)
	for id, rawConnection := range rawConnections {
		var connection Connection
		err = json.Unmarshal(rawConnection, &connection)
		if err != nil {
			println("Error unmarshalling connection")
			return nil
		}
		connections[id] = connection
	}

	return connections
}

// GetServer pulls the details of a specific server from the Guacamole API
func GetServer(serverID string, token string) Server {
	request, err := http.NewRequest("GET", guacURL+"/api/session/data/"+guacDatasource+"/connections/"+serverID+"?token="+token, nil)
	if err != nil {
		println("Error creating request")
		return Server{}
	}
	request.Header.Set("Accept", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		println("Error sending request")
		return Server{}
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			println("Error closing body")
			return
		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		println("Error reading body")
		return Server{}
	}

	var server Server
	err = json.Unmarshal(body, &server)
	if err != nil {
		println("Error unmarshalling json")
		return Server{}
	}
	return server
}

func ConnectionExists(queue []ConnectionDetail, connectionDetail ConnectionDetail) bool {
	for _, item := range queue {
		if item == connectionDetail {
			return true
		}
	}
	return false
}

func QueueContains(connections map[string]Connection, connectionDetail ConnectionDetail) bool {
	for _, connection := range connections {
		if connection.ConnectionIdentifier == connectionDetail.ConnectionID && connection.Username == connectionDetail.Username {
			return true
		}
	}
	return false
}
