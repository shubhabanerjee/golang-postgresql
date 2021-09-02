package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type userDetails struct {
	age        int
	first_name string
	last_name  string
	email      string
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

	data := getDataById(3, db)
	fmt.Println(data)
	// var u1 userDetails
	// u1.age = 20
	// u1.email = "bapan@gmail.com"
	// u1.first_name = "bapan"
	// u1.last_name = "banerjee"
	// data := setDataInDB(db, u1)
	// fmt.Println(data)
}

func setDataInDB(db *sql.DB, ud userDetails) int {
	sqlStatement := `
	  INSERT INTO users (age, email, first_name, last_name)
	  VALUES ($1, $2, $3, $4)
	  RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, ud.age, ud.email, ud.first_name, ud.last_name).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
	return id
}

//Get data Query
func getDataById(id int, db *sql.DB) userDetails {

	getDataQuery := `
	SELECT age, email, first_name, last_name FROM users WHERE id=$1;
	`
	var email, first_name, last_name string
	var age int
	data := db.QueryRow(getDataQuery, id)
	// fmt.Println(data)
	data.Scan(&age, &email, &first_name, &last_name)
	fmt.Println(age, email, first_name, last_name)
	var u1 userDetails
	u1.age = age
	u1.email = email
	u1.first_name = first_name
	u1.last_name = last_name
	return u1
}
