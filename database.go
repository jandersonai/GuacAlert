package main

import (
	"database/sql"
	"log"
)

func connectDB() {
	var err error
	db, err = sql.Open("postgres", "host="+dbIP+" port="+dbPort+" user="+dbUser+" password="+dbPass+" dbname="+dbName+" sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
