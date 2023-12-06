package groupModel

import (
	repository "be/src/models"
	"time"
)

type Group struct {
	ID         int                 `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string              `json:"name"`
	CreatedAt  time.Time           `json:"createdAt"`
	UpdatedAt  time.Time           `json:"updatedAt"`
	Items      []repository.Item   `gorm:"foreignKey:GroupId" json:"items"`
	AccountId  int                 `json:"accountId"`
	CategoryId int                 `json:"categoryId"`
	Category   repository.Category `json:"category"`
}
