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

}
// func NewRepository(db *gorm.DB) *Repository {
// 	return &Repository{db: db}
// }