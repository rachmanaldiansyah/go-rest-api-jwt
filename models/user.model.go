package models

type User struct {
	ID       int    `gorm:"primaryKey"`
	Name     string `gorm:"varchar(100)"`
	Email    string `gorm:"varchar(100)"`
	Password string `gorm:"varchar(100)"`
}

type Register struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}
