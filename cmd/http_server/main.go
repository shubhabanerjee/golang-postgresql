package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iamshubha/golang-postgresql/pkg/user"
	_ "github.com/lib/pq"
)

type userDetails struct {
	id         int    `json:"id"`
	age        int    `json:"age"`
	first_name string `json:"first_name"`
	last_name  string `json:"last_name"`
	email      string `json:"email"`
}

const (
	hostname     = "localhost"
	host_port    = 5432
	username     = "postgres"
	password     = "1234"
	databasename = "postgres"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		hostname, host_port, username, password, databasename)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var u1 userDetails
	u1.first_name = "Bapan"
	u1.last_name = "Banerjee"
	//user.GetDataById(db, 3)
	router := mux.NewRouter()

	// router.HandleFunc("/createUser", createUserDataInDB).Methods("POST")
	router.HandleFunc("/user/{id}", user.GetUserData).Methods("POST")
	http.ListenAndServe(":8080", router)
}
