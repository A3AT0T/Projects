package serverClient

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

var user []Person

func Webserver() {

	http.HandleFunc("/user", peopleHandler)
	http.HandleFunc("/check", checkHandler)

	log.Println("server starting in port 8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func peopleHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUser(w, r)
	case http.MethodPost:
		postPerson(w, r)
	default:
		http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
	}
}
func getUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(user)
	fmt.Fprint(w, "get people %v ", user)
}

func postPerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user = append(user, person)
	fmt.Fprint(w, "post new person: %v", person)
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Web-server OK")
}
