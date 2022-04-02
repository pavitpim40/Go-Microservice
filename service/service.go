package service

import (
	"fmt"
	"golang-project/repository"
)



type Servicer interface {
	Login(req Request) Response 
}


type Dber interface {
	LoginDB(user,pass string) (repository.TblLogin, error)
}
type Service struct {
	db Dber
}

func NewService(db Dber) *Service{
	return &Service{db :db}
}



func (s *Service)Login(req Request) Response{
	data,_ := s.db.LoginDB(req.Username,req.Password)
	fmt.Printf("data in service is %v",data)
	return Response{
		Fullname: data.Name,
		Birthday: data.Birthday,
		Age: data.Age,
	}
} 
