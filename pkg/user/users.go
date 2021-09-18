package user

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/iamshubha/golang-postgresql/pkg/model"
	"github.com/iamshubha/golang-postgresql/pkg/util"
	"golang.org/x/crypto/bcrypt"
)

type DB *sql.DB

func Signup(w http.ResponseWriter, r *http.Request) {
	db := util.GetDB()
	defer db.Close()
	creds := &model.Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
		return

	}
	if creds.Password == "" || creds.Username == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid request",
		})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8) //bcrypt.GenerateFromPassword([]byte(creds.Password), 8)
	if err != nil {
		log.Fatal(err)
	}
	sqlQuery := `
	INSERT INTO userlogin (username, password)
	VALUES ($1, $2)
	RETURNING id;
	`
	uid := model.UserSignupResponse{}
	err = db.QueryRow(sqlQuery, creds.Username, string(hashedPassword)).Scan(&uid.Uid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
	uid.Message = "Signup Success"

	json.NewEncoder(w).Encode(uid)

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	//Get DataBase
	db := util.GetDB()
	defer db.Close()

	// Get request body in creds
	creds := &model.Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)

	if err != nil {
		log.Println(err)
		return
	}
	if creds.Password == "" || creds.Username == "" {
		log.Println("noting")
		return
	}
	//Get response from postgresql database
	storedCreds := &model.Credentials{}
	uid := 0
	//getting data from database
	err = db.QueryRow("SELECT password, id FROM userlogin WHERE username=$1", creds.Username).Scan(&storedCreds.Password, &uid)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
			w.WriteHeader(404)
			json.NewEncoder(w).Encode(map[string]string{
				"Message": "No User Found !",
			})
			return
		}
		log.Println("no data found")
		log.Println(err)
		return

	}

	//Compare password
	err = bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(creds.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		d := struct {
			Message string `json:"message"`
		}{
			Message: "Wrong Password!",
		}
		json.NewEncoder(w).Encode(d)
		return
	}

	d := model.UserSignupResponse{}
	d.Uid = uid
	d.Message = "Login Success"
	//response back
	json.NewEncoder(w).Encode(d)

}
