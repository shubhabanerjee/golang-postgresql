package model

type GoleCreateMOdel struct {
	Userid int    `json:"userid"`
	Workon string `json:"workon"`
}

type IdAndUserid struct {
	Id     int `json:"id"`
	UserId int `json:"userid"`
}
