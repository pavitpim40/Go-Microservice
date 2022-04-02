package service

import "time"

type Request struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Fullname string `json:"fullname"`
	Birthday time.Time `json:"birthday"`
	Age int64 `json:"age"`
}