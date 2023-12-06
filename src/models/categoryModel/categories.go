package categoryModel

import (
	repository "be/src/models"
	"time"
)

type Category struct {
	ID        int               `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string            `json:"name"`
	CreatedAt time.Time         `json:"createdAt"`
	UpdatedAt time.Time         `json:"updatedAt"`
	Items     []repository.Item `gorm:"foreignKey:CategoryId" json:"items"`
	Groups    []repository.Group `gorm:"foreignKey:CategoryId" json:"groups"`
}
