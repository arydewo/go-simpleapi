package main
import (
    "database/sql"
)


type Users struct {
	id        int `json:"id"`
	FirstName sql.NullString `json:"firstName"`
	LastName  sql.NullString `json:"lastName"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Users
}
