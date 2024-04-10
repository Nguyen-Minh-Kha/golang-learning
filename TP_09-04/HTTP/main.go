package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type User struct {
	UserID   int
	Login    string `json:"userName"`
	Password string
}

var USERS map[string]User

// Function to populate the global variable USERS from a JSON file 
func populateUsers() map[string]User {
	content, err := os.ReadFile("./users.json")
	if err != nil {
		fmt.Println(err)
	}

	var users []User
	error := json.Unmarshal(content, &users)
	if error != nil {
		fmt.Println(err)
	}

	mapUsers := make(map[string]User)
	for _, user := range users {
		mapUsers["id"+strconv.Itoa((user.UserID))] = user
	}

	return mapUsers
}

func sendResponse(id string, w http.ResponseWriter) {
	if user, ok := USERS[id]; ok {
		response, err := json.Marshal(user)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set(
			"Content-Type",
			"application/json; charset=utf-8",
		)
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	} else {
		w.WriteHeader((http.StatusNotFound))
	}
}

// Using Handle method
func (u *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
	id := r.FormValue("id")
	sendResponse(id, w)
}

// Using handleFunc method
func handleFuncMethod () {
	defaultHandler := func(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("id")
		sendResponse(id, w)
	}
	http.HandleFunc("/handlef", defaultHandler)
}

// Function to serve the HTTP server
func serve() {
	handleFuncMethod()
	
	http.Handle("/handle", &User{})

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	USERS = populateUsers()

	serve()
}
