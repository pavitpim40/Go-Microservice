package repository
type TblLogin struct {
	Id       string `gorm:"column:id" json:"id"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Birthday string `gorm:"column:birthday" json:"birthday"`
	Age      string `gorm:"column:age" json:"age"`
	Name     string `gorm:"column:name" json:"name"`
   }
   
func (TblLogin) TableName() string {
	return "tbl_login"
}