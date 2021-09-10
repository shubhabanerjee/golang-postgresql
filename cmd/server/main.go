package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iamshubha/golang-postgresql/pkg/user"
	"github.com/iamshubha/golang-postgresql/pkg/util"
	_ "github.com/lib/pq"
)

func main() {
	db := util.GetDB()
	util.InitDB(db)
	defer db.Close()
	router := mux.NewRouter()
	router.HandleFunc("/user", user.GetUserData).Methods("POST")
	router.HandleFunc("/userCreate", user.CreateUser).Methods("POST")
	router.HandleFunc("/login", user.LoginHandler).Methods("POST")
	router.HandleFunc("/signup", user.Signup).Methods("POST")
	http.ListenAndServe(":8080", router)
}
