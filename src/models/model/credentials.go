package credentialModel

import (
	"gorm.io/gorm"
)

type Credential struct {
	gorm.Model
	ID        int    `gorm:"primaryKey;autoIncrement"`
	Key       string `gorm:"not null"`
	Email     string `gorm:"not null"`
	Username  string `gorm:"not null"`
	Password  string `gorm:"not null"`
	AccountId int
}
