package main

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/webhook", WebhookHandler).Methods("POST")
	router.HandleFunc("/status", StatusHandler).Methods("POST")
	router.HandleFunc("/healthcheck", HealthCheckHandler).Methods("GET")
	return router
}

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	// Implement the function to parse the body from Google Workspace Chat, parse the message, and send the appropriate message to the chat room

}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	// Implement the function to check the status of a machine and respond with the user who is currently controlling the machine
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	write, err := w.Write([]byte("OK"))
	if err != nil {
		return
	}
	if write != 2 {
		return
	}
}
