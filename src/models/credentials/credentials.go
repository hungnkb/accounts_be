package model

type Credential struct {
	Id       int `gorm:"primaryKey;autoIncrement"`
	Key      string
	Email    string
	Username string
}
