package model

type UserSignupResponse struct {
	Message string `json:"message"`
	Uid     int    `json:"id"`
}

type UserDetailsResponse struct {
	Id         int    `json:"id"`
	Age        int    `json:"age"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
}
type UserDetailsResponseGetFromUser struct {
	Age        int    `json:"age"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
}

type User struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type ReturnMessage struct {
	Message string                `json:"message"`
	Data    []UserDetailsResponse `json:"data"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
