package model

type Account struct {
	Id       int `gorm:"primaryKey;autoIncrement"`
	Name     string
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
}
