package models

type User struct {
	ID        uint     `gorm:"primary_key" json:"id"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Password  string   `json:"-"`
	Devices   []Device `gorm:"many2many:user_devices;" json:"-"`
	PushToken string   `json:"pushToken"`
}
