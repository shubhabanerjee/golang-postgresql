package model

type TaskFormat struct {
	Uid    int    `json:"id"`
	Body   string `json:"body"`
	Title  string `json:"title"`
	Bucket string `json:"bucket"`
}

type GetTaskData struct {
	Body  string `json:"body"`
	Title string `json:"title"`
}

type TaskUpdateFormat struct {
	Id     int    `json:"id"`
	Uid    int    `json:"userid"`
	Body   string `json:"body"`
	Title  string `json:"title"`
	Bucket string `json:"bucket"`
}
