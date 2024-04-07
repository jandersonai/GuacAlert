package main

import (
	"database/sql"
	"os"
)

var (
	listenPort = os.Getenv("LISTEN_PORT")
	dbIP       = os.Getenv("DB_IP")
	dbPort     = os.Getenv("DB_PORT")
	dbUser     = os.Getenv("DB_USER")
	dbPass     = os.Getenv("DB_PASS")
	dbName     = os.Getenv("DB_NAME")
	chatURL    = os.Getenv("CHAT_HOOK")
	db         *sql.DB
	queue      []string
)
