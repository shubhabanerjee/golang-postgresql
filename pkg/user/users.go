package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
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
		log.Fatal(err)
		return
	}
	//Get response from postgresql database
	storedCreds := &model.Credentials{}
	uid := 0
	//getting data from database
	err = db.QueryRow("SELECT password, id FROM userlogin WHERE username=$1", creds.Username).Scan(&storedCreds.Password, &uid)

	if err != nil {

		if err == sql.ErrNoRows {
			log.Fatal(err)
			return
		}
		log.Fatal(err)
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

func GenerateJWT(email string, id int) (string, error) {
	var mySigningKey = []byte("secretkey")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func GetUserData(w http.ResponseWriter, r *http.Request) {
	var u model.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Hi! my first name is %s and surname is %s\n", u.Name, u.Surname)
	err = json.NewEncoder(w).Encode(u)
	if err != nil {
		log.Fatal(err)
		return
	}
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

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var uD model.UserDetailsResponseGetFromUser
	var rsp model.ReturnMessage
	err := json.NewDecoder(r.Body).Decode(&uD)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(uD)
	db := util.GetDB()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()
	fmt.Println(uD)
	response := SetDataInDB(db, uD)
	data := GetDataById(db, response)
	rsp.Message = "User Created Successfully"
	rsp.Data = []model.UserDetailsResponse{data}
	err = json.NewEncoder(w).Encode(rsp)
	if err != nil {
		log.Fatal(err)
		return
	}
}

//var posts []UserDetailsResponse
//
////Delete data Query
//func DeleteDataInDB(db *sql.DB, id int) bool {
//	sqlQuery := `
//	DELETE FROM users
//	WHERE id = $1;
//	`
//	_, err := db.Exec(sqlQuery, id)
//	if err != nil {
//		log.Fatal(err)
//
//	}
//	fmt.Println("delete success")
//	return true
//}
//
////Update data Query
//func UpdateDataInDB(db *sql.DB, id int, ud UserDetailsResponse) UserDetailsResponse {
//	sqlUpdate := `
//	UPDATE users
//	SET first_name = $2, last_name = $3
//	WHERE id = $1;
//	`
//	_, err := db.Exec(sqlUpdate, id, ud.first_name, ud.last_name)
//	if err != nil {
//		log.Fatal(err)
//	}
//	data := getDataById(db, id)
//	return data
//}
//
////Set data Query
func SetDataInDB(db *sql.DB, ud model.UserDetailsResponseGetFromUser) int {
	fmt.Println(ud)
	sqlStatement := `
	INSERT INTO users (age, email, first_name, last_name)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, ud.Age, ud.Email, ud.First_name, ud.Last_name).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
	return id
}

//
////Get data Query
func GetDataById(db *sql.DB, id int) model.UserDetailsResponse {
	var response model.UserDetailsResponse
	fmt.Println("GetDataById")
	getDataQuery := `
	SELECT age, email, first_name, last_name FROM users WHERE id=$1;
	`
	var email, first_name, last_name string
	var age int
	data := db.QueryRow(getDataQuery, id)
	// fmt.Println(data)
	response.Id = id
	data.Scan(&response.Age, &response.Email, &response.First_name, &response.Last_name)
	fmt.Println(age, email, first_name, last_name)

	return response
}

//
//// func createUserDataInDB(w http.ResponseWriter, r *http.Request) {
//// 	w.Header().Set("Content-Type", "application/json")
//// 	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
//// 	w.Header().Set("Access-Control-Allow-Origin", "*")
//// 	w.Header().Set("Access-Control-Allow-Methods", "POST")
//// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
//// 	var p UserDetailsResponse
//// 	json.NewDecoder(r.Body).Decode(&p)
//// 	fmt.Println(p.email)
//// 	// posts = append(posts, p)
//// 	json.NewEncoder(w).Encode(p)
//
//// }
//
////Update data Query
