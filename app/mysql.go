package main

import (
	"database/sql"
	"log"
	"os"
)

func connect() *sql.DB {
	db, err := sql.Open(os.Getenv("dbtype"), os.Getenv("dburl"))

	if err != nil {
		log.Fatal(err)
	}

	return db
}
