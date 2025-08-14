package service

import "time"

type Word struct {
	Title       string `json:"title"`
	Translation string `json:"translation"`
}

type Report struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Overview   string    `json:"overview"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
