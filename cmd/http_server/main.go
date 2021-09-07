package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iamshubha/golang-postgresql/pkg/user"
	_ "github.com/lib/pq"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user", user.GetUserData).Methods("POST")
	router.HandleFunc("/userCreate", user.CreateUser).Methods("POST")
	router.HandleFunc("/signup", user.Signup).Methods("POST")

	http.ListenAndServe(":8080", router)
}
