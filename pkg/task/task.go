package task

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/iamshubha/golang-postgresql/pkg/model"
	"github.com/iamshubha/golang-postgresql/pkg/util"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	db := util.GetDB()
	defer db.Close()
	creds := &model.TaskCreateFormat{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if creds.Uid == 0 || creds.Body == "" || creds.Title == "" || creds.Bucket == "" {
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Please send correct parameaters",
		})
		return
	}
	if err != nil {
		log.Print(err)
		return
	}
	sqlQuery := `
	INSERT INTO tasktable (userid, bucket, title, body, created_at, update_on)
	VALUES ($1,$2,$3,$4,$5,$6) 
	RETURNING id;
	`
	_, err = db.Exec(sqlQuery, creds.Uid, creds.Bucket, creds.Title, creds.Body, time.Now(), time.Now())
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Fail",
		})
		log.Print(err)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"message": "success",
	})

}

func GetTask(w http.ResponseWriter, r *http.Request) {

	urlData := mux.Vars(r)
	id, ok := urlData["id"]
	if !ok {
		log.Println(ok)
	}
	db := util.GetDB()
	defer db.Close()
	sqlQuery := `
	SELECT title, body FROM tasktable WHERE userid = $1;
	`

	dataRow, err := db.Query(sqlQuery, id)
	if err != nil {
		log.Println(err)
		return
	}
	defer dataRow.Close()
	data := make([]model.GetTaskData, 0)
	for dataRow.Next() {
		ddd := model.GetTaskData{}
		dataRow.Scan(&ddd.Title, &ddd.Body)
		data = append(data, ddd)
	}
	if len(data) == 0 {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "No data found on this user",
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "success",
		"data":    data,
	})
	fmt.Println(data)
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
