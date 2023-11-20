package accountModel

import (
	credential "be/src/models/model"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	ID          int `gorm:"primaryKey;autoIncrement"`
	Name        string
	Username    string                  `gorm:"unique;not null"`
	Email       string                  `gorm:"unique;not null"`
	Credentials []credential.Credential `gorm:"foreignKey:AccountId"`
}
