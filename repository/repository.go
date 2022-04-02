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
// func NewRepository(db *gorm.DB) *Repository {
// 	return &Repository{db: db}
// }