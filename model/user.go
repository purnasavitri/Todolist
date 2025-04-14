package model

type User struct {
	ID       uint   `gorm:"column:id_user;primaryKey;table:user" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"-"`
}

func (User) TableName() string {
	return "user"
}