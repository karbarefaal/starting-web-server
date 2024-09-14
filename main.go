package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type user struct {
	Name   string
	Family string
	Age    uint16
}

//=============================================================================================

func createUser(w http.ResponseWriter, r *http.Request, users *[]user) {
	name := r.FormValue("name")
	family := r.FormValue("family")

	*users = append(*users, user{
		Name:   name,
		Family: family,
		Age:    26,
	})
	fmt.Println(*users)
}

//=============================================================================================

func updateUser(w http.ResponseWriter, r *http.Request) {

}

// =============================================================================================
func deleteUser(w http.ResponseWriter, r *http.Request /* , users *[]user, string name   */) {
	// users = delete(users, name)
}

func scanName(name string) {

}

//=============================================================================================

func getUser(w http.ResponseWriter, r *http.Request, users *[]user) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(*users)
}

//=============================================================================================

func main() {

	// var users []user

	users := new([]user)

	// /users is endpoint for our url
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case "GET":
			getUser(w, r, users)
			// fmt.Println("Methoe : GET")
		case "POST":
			createUser(w, r, users)
			// fmt.Println("Methoe : POST")
		case "PUT":
			fmt.Println("Methoe : PUT")
		case "DELETE":
			deleteUser(w, r /*  , users, user  */)
			// fmt.Println("Methoe : DELETE")
		}
		fmt.Fprintf(w, "www.uncodev.com")
	})

	err := http.ListenAndServe(":8989", nil)
	if err != nil {
		fmt.Println("shit")
		log.Fatal(err)
	}
}
