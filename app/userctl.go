package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func returnAllUsers(w http.ResponseWriter, r *http.Request) {
	var users Users
	var arr_user []Users
	var response Response

	db := connect()
	defer db.Close()

	rows, err := db.Query("Select id,first_name as FirstName,last_name as LastName from users where first_name is not null and last_name is not null limit 0,5;")
	
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&users.id, &users.FirstName, &users.LastName); err != nil {
			log.Fatal(err.Error())

		} else {
			log.Print(users)
			arr_user = append(arr_user, users)
		}
	}
	log.Print(arr_user)
	response.Status = 1
	response.Message = "Success"
	response.Data = arr_user

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
