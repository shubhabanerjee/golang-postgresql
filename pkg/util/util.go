package util

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	hostname     = "localhost"
	host_port    = 5432
	username     = "postgres"
	password     = "1234"
	databasename = "postgres"
)

func GetDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		hostname, host_port, username, password, databasename)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}

func InitDB(db *sql.DB) {
	sqlQueryforlogintable := `
	CREATE TABLE IF NOT EXISTS userlogin (
		id SERIAL PRIMARY KEY,
		userName TEXT NOT NULL,
		password TEXT NOT NULL
	);	
	`
	data, err := db.Exec(sqlQueryforlogintable)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(data)

	sqlQueryfortasktable := `
	CREATE TABLE IF NOT EXISTS  tasktable (
		id SERIAL PRIMARY KEY,
	userid INTEGER NOT NULL,
	title TEXT NOT NULL,
	body TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL,
	update_on TIMESTAMP NOT NULL
	
	);
	`
	data, err = db.Exec(sqlQueryfortasktable)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(data)
}
