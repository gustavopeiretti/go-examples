package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HomePage!")
	fmt.Println("Endpoint homePage")
}

func handleView(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		keys := r.URL.Query()
		id := keys.Get("id") //Get return empty string if key not found
		if id == "" {
			log.Println("Url Param 'id' is missing")
			return
		}
		// simulate search by id and return the user for this id
		idInt, _ := strconv.ParseInt(id, 10, 64)
		user := User{
			Id:   idInt,
			Name: "gustavo",
		}
		//
		jsonUser, err := json.Marshal(user)
		if err != nil {
			log.Panic(err)
		}
		w.Write(jsonUser)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func handleSave(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		var user User
		err := decoder.Decode(&user)
		if err != nil {
			panic(err)
		}
		// simulate save user
		user.Id = 1234
		//
		// return the saved user
		jsonUser, err := json.Marshal(user)
		if err != nil {
			log.Panic(err)
		}
		w.Write(jsonUser)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/view", handleView)
	http.HandleFunc("/save", handleSave)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
