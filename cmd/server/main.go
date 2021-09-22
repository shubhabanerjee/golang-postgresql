package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iamshubha/golang-postgresql/pkg/gole"
	"github.com/iamshubha/golang-postgresql/pkg/task"
	"github.com/iamshubha/golang-postgresql/pkg/user"
	"github.com/iamshubha/golang-postgresql/pkg/util"
	_ "github.com/lib/pq"
)

func main() {
	db := util.GetDB()
	util.InitDB(db)
	defer db.Close()
	router := mux.NewRouter()
	//TODO:forget, delete note
	router.HandleFunc("/deleteNote", task.DeleteTask).Methods(http.MethodPost)
	router.HandleFunc("/createNote", task.CreateTask).Methods(http.MethodPost)
	router.HandleFunc("/updateTask", task.UpdateTask).Methods(http.MethodPost)
	router.HandleFunc("/notes/userid:{id}", task.GetTask).Methods(http.MethodGet)
	router.HandleFunc("/GetNoteFromBucket/id:{id}/bucket:{bucket}", task.GetTaskFromBucket).Methods(http.MethodGet)
	router.HandleFunc("/login", user.LoginHandler).Methods(http.MethodPost)
	router.HandleFunc("/signup", user.Signup).Methods(http.MethodPost)
	router.HandleFunc("/forgetPassword", user.ForgetPassword).Methods(http.MethodPost)
	router.HandleFunc("/startWorking", gole.StartWorking).Methods(http.MethodPost)
	router.HandleFunc("/stopWorking", gole.StopWorking).Methods(http.MethodPost)
	http.ListenAndServe(":8080", router)

}

//Get filtered query
//SELECT * FROM tasktable WHERE userid = 1 AND title = 'n';
// func Tasks(w http.ResponseWriter, r *http.Request) {
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	fmt.Println(body)
// 	urlparams := mux.Vars(r)
// 	id, ok := urlparams["id"]
// 	if !ok {
// 		log.Println(ok)
// 	}
// }

// func GenerateJWT(email string, id int) (string, error) {
// 	var mySigningKey = []byte("secretkey")
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["authorized"] = true
// 	claims["email"] = email
// 	claims["id"] = id
// 	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
// 	tokenString, err := token.SignedString(mySigningKey)
// 	if err != nil {
// 		fmt.Errorf("Something Went Wrong: %s", err.Error())
// 		return "", err
// 	}
// 	return tokenString, nil
// }
// func getdata() {
// 	db := util.GetDB()
// 	defer db.Close()
// 	// "SELECT password, id FROM userlogin WHERE username=$1",
// 	q := `
// 	SELECT * FROM tasktable WHERE userid = 2;
// 	`
// 	// d := model.TaskFormat{}
// 	type d struct {
// 		userid     int    `json:"userid"`
// 		created_at string `json:"created_at"`
// 		update_on  string `json:"update_on"`
// 		Uid        int    `json:"id"`
// 		Body       string `json:"body"`
// 		Title      string `json:"title"`
// 	}
// 	dd := d{}
// 	data := db.QueryRow(q).Scan(&dd.userid, &dd.Title, &dd.Body, &dd.created_at, &dd.update_on, &dd.Uid)
// 	fmt.Println(data)
// 	fmt.Println(dd)
// 	log.Print("------------------------------------------")
// 	fmt.Println(dd.created_at)
// 	fmt.Println(dd)
// 	log.Print("------------------------------------------++++")
// 	ddddataRow, err := db.Query(q)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer ddddataRow.Close()
// 	snbs := make([]d, 0)
// 	for ddddataRow.Next() {
// 		ddd := d{}
// 		ddddataRow.Scan(&ddd.userid, &ddd.Title, &ddd.Body, &ddd.created_at, &ddd.update_on, &ddd.Uid)
// 		snbs = append(snbs, ddd)
// 	}
// 	fmt.Println(snbs)
// }
