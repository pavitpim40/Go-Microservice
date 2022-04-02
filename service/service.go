package service



type Servicer interface {
	Login(req Request) Response 
}

func (s *Service)Login(req Request) Response{
	return Response{}
} 

type Service struct {

}

func NewService() *Service{
	return &Service{}
}
