package main

import (
	"net/http"
	"strings"
)

// SendMessage Send a message to Google Workspace Chat
func SendMessage(message string) {
	messageBody := strings.NewReader(`{"text":"` + message + `"}`)
	request, err := http.NewRequest("POST", chatURL, messageBody)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return
	}
	defer func() {
		err := response.Body.Close()
		if err != nil {
			return
		}
	}()
}
