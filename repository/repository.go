package repository

import (
	"github.com/jinzhu/gorm"
)

type Repository struct {
	db *gorm.DB
}


func New(db *gorm.DB) *Repository {
	return &Repository {db:db}
}

func (db *Repository)LoginDB(user,pass string) (TblLogin, error) {
	var result TblLogin
	db.db.Where("username = ? AND password = ?", user, pass).First(&result)
	return result, nil
}

func (db *Repository) Register(req TblLogin) error {

	if err :=	db.db.Create(&req).Error; err != nil {
		return err
	}
	return nil
	// reqDB := repository.TblLogin{
	// 	Username : req.Username,
	// 	Password : req.Password,
	// 	BirthDay : req.Birthday,
	// 	Age : req.Age,
	// }
	// if err := s.db.Register(reqDB);err!= nil {
	// 	fmt.Printf("error in register is %v",err)
	// }
	// return req, nil
}
// func NewRepository(db *gorm.DB) *Repository {
// 	return &Repository{db: db}
// }