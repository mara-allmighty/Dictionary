package service

type Word struct {
	Title       string `json:"title"`
	Translation string `json:"translation"`
}

type Report struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Overview string `json:"overview"`
}
