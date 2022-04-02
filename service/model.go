package service

type Request struct {
	Username string `json:"username"`
	Password string `json:"password"`
}


type Response struct {
	Fullname string `json:"fullname"`
	Birthday string `json:"birthday"`
	Age int64 `json:"age"`
}


type RequestReg struct {
	Username string    `json:"username"`
	Password string    `json:"password"`
	Birthday string `json:"birthday"`
	Age      int64     `json:"age"`
   }

type ResponseReg struct {
}