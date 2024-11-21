package models

type User struct {
	ID       int    `gorm:"primaryKey"`
	Name     string `gorm:"varchar(100)"`
	Email    string `gorm:"varchar(100)"`
	Password string `gorm:"varchar(100)"`
}
