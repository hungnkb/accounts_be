package itemModel

import (
	repository "be/src/models"
	"time"
)

type Item struct {
	ID        int              `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string           `json:"name"`
	Username  string           `json:"username"`
	Email     string           `json:"email"`
	Password  string           `json:"password"`
	CreatedAt time.Time        `json:"createdAt"`
	UpdatedAt time.Time        `json:"updatedAt"`
	AccountId *int             `json:"accountId"`
	GroupId   *int             `json:"groupId"`
	Group     repository.Group `json:"group"`
}
