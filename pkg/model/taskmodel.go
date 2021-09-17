package model

type TaskCreateFormat struct {
	Uid   int    `json:"id"`
	Body  string `json:"body"`
	Title string `json:"title"`
	// Bucket string `json:"bucket"`
}

type GetTaskData struct {
	Body  string `json:"body"`
	Title string `json:"title"`
}
