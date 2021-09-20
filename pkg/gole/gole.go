package gole

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/iamshubha/golang-postgresql/pkg/model"
	"github.com/iamshubha/golang-postgresql/pkg/util"
)

func StartWorking(w http.ResponseWriter, r *http.Request) {
	golemodel := model.GoleCreateMOdel{}
	db := util.GetDB()
	defer db.Close()
	err := json.NewDecoder(r.Body).Decode(&golemodel)
	if err != nil {
		fmt.Println(err)
		return
	}
	if golemodel.Userid == 0 || golemodel.Workon == "" {
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Please send correct parameaters",
		})
		return
	}
	sqlQuery := `
	INSERT INTO goletable (userid, workon, starttime)
	VALUES ($1,$2,$3)  
	RETURNING id;
	`
	_, err = db.Exec(sqlQuery, golemodel.Userid, golemodel.Workon, time.Now())
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

func StopWorking(w http.ResponseWriter, r *http.Request) {

	type stopWork struct {
		Id     int `json:"id"`
		UserId int `json:"userid"`
	}

	golemodel := stopWork{}
	db := util.GetDB()
	defer db.Close()
	err := json.NewDecoder(r.Body).Decode(&golemodel)
	if err != nil {
		fmt.Println(err)
		return
	}
	if golemodel.UserId == 0 || golemodel.Id == 0 {
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Please send correct parameaters",
		})
		return
	}
	sqlQuery := `
	UPDATE goletable SET stoptime = $3 
	WHERE userid = $1 AND id = $2;
	`
	_, err = db.Exec(sqlQuery, golemodel.UserId, golemodel.Id, time.Now())
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
