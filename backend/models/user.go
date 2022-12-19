package models

type User struct {
	Id       int64  `gorm:"primary_key" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}
