package task

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/iamshubha/golang-postgresql/pkg/model"
	"github.com/iamshubha/golang-postgresql/pkg/util"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	db := util.GetDB()
	defer db.Close()
	creds := &model.TaskCreateFormat{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		log.Print(err)
		return
	}
	sqlQuery := `
	INSERT INTO tasktable (userid, title, body, created_at, update_on)
	VALUES ($1,$2,$3,$4,$5)
	RETURNING id;
	`
	ok, err := db.Exec(sqlQuery, creds.Uid, creds.Title, creds.Body, time.Now(), time.Now())
	if err != nil {
		log.Print(err)
		return
	}
	fmt.Println(ok)

}

/*
	tasktable =>

	id SERIAL PRIMARY KEY,
	userid INTEGER NOT NULL,
	title TEXT NOT NULL,
	body TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL,
	update_on TIMESTAMP NOT NULL
*/
