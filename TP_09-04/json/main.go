package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Login string `json:"userName"`
	Password string
}

// Serialize the user to JSON
func ex1 () {	
	user := User{Login: "Paul", Password: "pass123"}
	b, err := json.Marshal(user)
	if err!= nil {
        fmt.Println(err)
    }
	fmt.Println(string(b))
}

// Deserialize the user from JSON
func ex2 () {
	content, err := os.ReadFile("./users.json")
	if err!= nil {
        fmt.Println(err)
    }

	var users []User
	
	error := json.Unmarshal(content, &users)
	
	if error!= nil {
        fmt.Println(err)
    }
	fmt.Printf("%+v", users)
}

func main() {
    ex1()

	ex2()
	
}