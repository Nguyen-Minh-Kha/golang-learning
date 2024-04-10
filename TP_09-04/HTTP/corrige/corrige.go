package main

import (
  "encoding/json"
  "fmt"
  "net/http"
  "os"
)

var userMap map[string]User

type User struct {
  UserID   string
  Login    string `json:"userName"`
  Password string
}

func main() {
  fillDatabase()
  http.Handle("/", &User{})
  http.ListenAndServe("localhost:4000", nil)
}

func fillDatabase() {
  file, _ := os.ReadFile("users.json")
  var users []User
  json.Unmarshal(file, &users)

  userMap = make(map[string]User)
  for _, user := range users {
    userMap[user.UserID] = user
  }
}

func (s *User) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
  id := req.FormValue("id")
  idFound := searchId(id)
  sendResponse(idFound, rw, id)
}

func sendResponse(idFound bool, rw http.ResponseWriter, id string) {
  if idFound {
    rw.Header().Set("Content-Type", "application/json; charset=utf-8")
    bArray, _ := json.Marshal(userMap[id])
    rw.Write(bArray)
    rw.WriteHeader(http.StatusOK)
  } else {
    rw.WriteHeader(http.StatusNotFound)
  }
}

func searchId(id string) bool {
  var idFound bool
  for index := range userMap {
    if index == id {
      fmt.Println("Id trouvé !")
      idFound = true
      break
    }
  }
  return idFound
}