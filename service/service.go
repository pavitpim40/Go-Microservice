package service

import (
	"fmt"
	"golang-project/repository"
)



type Servicer interface {
	Login(req Request) Response 
	Register(req RequestReg) ResponseReg
}


type Dber interface {
	Register(req repository.TblLogin) error
	LoginDB(user,pass string) (repository.TblLogin, error)
}
type Service struct {
	db Dber
}

func NewService(db Dber) *Service{
	return &Service{db :db}
}

func (s *Service) Register (req RequestReg) ResponseReg{
		reqDB := repository.TblLogin{
		Username : req.Username,
		Password : req.Password,
		Birthday : req.Birthday,
		Age : req.Age,
	}
	if err := s.db.Register(reqDB);err!= nil {
		fmt.Printf("error in register is %v",err)
	}

	return ResponseReg{
	}
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
