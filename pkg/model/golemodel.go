package model

type GoleCreateMOdel struct {
	Userid int    `json:"userid"`
	Workon string `json:"workon"`
}

type IdAndUserid struct {
	Id     int `json:"id"`
	UserId int `json:"userid"`
}

type GoleDetails struct {
	Id        int    `json:"id"`
	Userid    int    `json:"userid"`
	Workon    string `json:"workon"`
	Starttime string `json:"starttime"`
	Stoptime  string `json:"stoptime"`
	Total     string `json:"total"`
}
