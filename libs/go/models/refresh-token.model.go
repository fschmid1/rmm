package models

type RefreshToken struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Token     string `json:"token"`
	UserID    uint   `json:"userId"`
	ExpiresAt int64  `json:"expiresAt"`
}
