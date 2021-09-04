package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
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
	getDataById(db, 3)
	router := mux.NewRouter()

	// router.HandleFunc("/createUser", createUserDataInDB).Methods("POST")
	router.HandleFunc("/user/{id}/", getUserData).Methods("GET")
	http.ListenAndServe(":8080", router)
}

func getUserData(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Body: ", string(body))
	urlParams := mux.Vars(r)
	id, ok := userParams["id"]
	if !ok {
		fmt.Println("ID not found")
		return
	}

	fmt.Println("print id:", id)
	//// get the ID of the post from the route parameter
	//var idParam string = mux.Vars(r)["id"]
	//id, err := strconv.Atoi(idParam)
	//if err != nil {
	//	// there was an error
	//	w.WriteHeader(400)
	//	w.Write([]byte("ID could not be converted to integer"))
	//	return
	//}

	//// error checking
	//if id >= len(posts) {
	//	w.WriteHeader(404)
	//	w.Write([]byte("No post found with specified ID"))
	//	return
	//}

	//post := posts[id]

	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(post)
}

var posts []userDetails

//Delete data Query
func deleteDataInDB(db *sql.DB, id int) bool {
	sqlQuery := `
	DELETE FROM users
	WHERE id = $1;
	`
	_, err := db.Exec(sqlQuery, id)
	if err != nil {
		panic(err)

	}
	fmt.Println("delete success")
	return true
}

//Update data Query
func updateDataInDB(db *sql.DB, id int, ud userDetails) userDetails {
	sqlUpdate := `
	UPDATE users
	SET first_name = $2, last_name = $3
	WHERE id = $1;
	`
	_, err := db.Exec(sqlUpdate, id, ud.first_name, ud.last_name)
	if err != nil {
		panic(err)
	}
	data := getDataById(db, id)
	return data
}

//Set data Query
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
func getDataById(db *sql.DB, id int) userDetails {

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

// func createUserDataInDB(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "POST")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 	var p userDetails
// 	json.NewDecoder(r.Body).Decode(&p)
// 	fmt.Println(p.email)
// 	// posts = append(posts, p)
// 	json.NewEncoder(w).Encode(p)

// }

//Update data Query
