package main

import (
	"io"
	"net/http"
)

// GenerateToken Generate a token for the user to use to authenticate with the Guacamole API
func GenerateToken() string {
	// Query the Guacamole API using username and password to generate a token
	request, err := http.NewRequest("POST", guacURL+"/api/tokens", nil)
	if err != nil {
		return ""
	}
	request.SetBasicAuth(guacUser, guacPass)
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
	return string(token)
}
