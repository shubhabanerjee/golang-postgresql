package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func GetUserData(w http.ResponseWriter, r *http.Request) {
	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Hi! my first name is %s and surname is %s\n", u.Name, u.Surname)
	urlParams := mux.Vars(r)
	id, ok := urlParams["id"]
	if !ok {
		fmt.Println("ID not found")
		return
	}
	fmt.Println("print id:", id)
	//fmt.Println(urlParams["country"])
	country := r.URL.Query().Get("country")
	fmt.Println("Country: ", country)

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

//var posts []userDetails
//
////Delete data Query
//func DeleteDataInDB(db *sql.DB, id int) bool {
//	sqlQuery := `
//	DELETE FROM users
//	WHERE id = $1;
//	`
//	_, err := db.Exec(sqlQuery, id)
//	if err != nil {
//		panic(err)
//
//	}
//	fmt.Println("delete success")
//	return true
//}
//
////Update data Query
//func UpdateDataInDB(db *sql.DB, id int, ud userDetails) userDetails {
//	sqlUpdate := `
//	UPDATE users
//	SET first_name = $2, last_name = $3
//	WHERE id = $1;
//	`
//	_, err := db.Exec(sqlUpdate, id, ud.first_name, ud.last_name)
//	if err != nil {
//		panic(err)
//	}
//	data := getDataById(db, id)
//	return data
//}
//
////Set data Query
//func SetDataInDB(db *sql.DB, ud userDetails) int {
//	sqlStatement := `
//	INSERT INTO users (age, email, first_name, last_name)
//	VALUES ($1, $2, $3, $4)
//	RETURNING id`
//	id := 0
//	err := db.QueryRow(sqlStatement, ud.age, ud.email, ud.first_name, ud.last_name).Scan(&id)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("New record ID is:", id)
//	return id
//}
//
////Get data Query
//func GetDataById(db *sql.DB, id int) userDetails {
//
//	getDataQuery := `
//	SELECT age, email, first_name, last_name FROM users WHERE id=$1;
//	`
//	var email, first_name, last_name string
//	var age int
//	data := db.QueryRow(getDataQuery, id)
//	// fmt.Println(data)
//	data.Scan(&age, &email, &first_name, &last_name)
//	fmt.Println(age, email, first_name, last_name)
//	var u1 userDetails
//	u1.age = age
//	u1.email = email
//	u1.first_name = first_name
//	u1.last_name = last_name
//	return u1
//}
//
//// func createUserDataInDB(w http.ResponseWriter, r *http.Request) {
//// 	w.Header().Set("Content-Type", "application/json")
//// 	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
//// 	w.Header().Set("Access-Control-Allow-Origin", "*")
//// 	w.Header().Set("Access-Control-Allow-Methods", "POST")
//// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
//// 	var p userDetails
//// 	json.NewDecoder(r.Body).Decode(&p)
//// 	fmt.Println(p.email)
//// 	// posts = append(posts, p)
//// 	json.NewEncoder(w).Encode(p)
//
//// }
//
////Update data Query
