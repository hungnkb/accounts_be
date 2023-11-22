package credentialModel

import (
	"time"
)

type Credential struct {
	// gorm.Model
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Key       string    `gorm:"not null" json:"key"`
	Email     string    `gorm:"not null" json:"email"`
	Username  string    `gorm:"not null" json:"username"`
	Password  string    `gorm:"not null" json:"password"`
	AccountId int       `gorm:"not null" json:"accountId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
