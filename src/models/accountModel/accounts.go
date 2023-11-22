package accountModel

import (
	credential "be/src/models/model"
	"time"
)

type Account struct {
	// gorm.Model
	ID          int                     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string                  `json:"name"`
	Username    string                  `gorm:"unique;not null" json:"username"`
	Email       string                  `gorm:"unique;not null" json:"email"`
	CreatedAt   time.Time               `json:"createdAt"`
	UpdatedAt   time.Time               `json:"updatedAt"`
	Credentials []credential.Credential `gorm:"foreignKey:AccountId" json:"credentials"`
}
