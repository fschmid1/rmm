package models

type User struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	PushToken string `json:"pushToken"`
}
