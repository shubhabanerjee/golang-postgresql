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
	creds := &model.TaskFormat{}
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

func GetTaskFromBucket(w http.ResponseWriter, r *http.Request) {

	urlData := mux.Vars(r)
	id, ok := urlData["id"]
	if !ok {
		log.Println(ok)
		w.WriteHeader(400)
		return
	}
	bucket, ok := urlData["bucket"]
	if !ok {
		log.Println(ok)
		w.WriteHeader(400)
		return
	}
	db := util.GetDB()
	defer db.Close()
	sqlQuery := `
	SELECT title, body FROM tasktable WHERE userid = $1 AND bucket = $2;
	`

	dataRow, err := db.Query(sqlQuery, id, bucket)
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

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	dataModel := model.TaskUpdateFormat{}
	db := util.GetDB()
	defer db.Close()
	err := json.NewDecoder(r.Body).Decode(&dataModel)
	if err != nil {
		w.WriteHeader(400)
		log.Println(err)
		return
	}
	fmt.Println(dataModel)
	if dataModel.Body == "" || dataModel.Bucket == "" || dataModel.Title == "" || dataModel.Uid == 0 || dataModel.Id == 0 {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "parameter cann't be empty",
		})
		return
	}
	sqlQuery := `
	UPDATE tasktable SET title = $1, body = $2, update_on = $3 WHERE userid = $4 AND id = $5;
	`
	_, err = db.Query(sqlQuery, dataModel.Title, dataModel.Body, time.Now(), dataModel.Uid, dataModel.Id)

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

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	sqlQuery := `
	DELETE FROM tasktable
	WHERE  userid = $1 AND id = $2;
	`
	db := util.GetDB()
	defer db.Close()
	type userDetails struct {
		Id     int `json:"id"`
		Userid int `json:"userid"`
	}
	d := userDetails{}
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = db.Exec(sqlQuery, d.Userid, d.Id)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Fail",
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Succes",
	})

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
