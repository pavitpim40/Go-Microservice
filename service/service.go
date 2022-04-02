package service



type Servicer interface {
	Login(req Request) Response 
}

type Service struct {

}

func NewService() *Service{
	return &Service{}
}



func (s *Service)Login(req Request) Response{
	return Response{}
} 
