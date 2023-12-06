package repository

import "time"

type Item struct {
	ID         int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string    `json:"name"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	GroupId    *int      `json:"groupId"`
	AccountId  int       `json:"accountId"`
	CategoryId *int      `json:"categoryId"`
}

type Category struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Items     []Item    `gorm:"foreignKey:CategoryId" json:"items"`
	Groups    []Group   `gorm:"foreignKey:CategoryId" json:"groups"`
}

type Group struct {
	ID         int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	AccountId  int       `json:"accountId"`
	CategoryId int       `json:"categoryId"`
	Category   Category  `json:"category"`
	Items      []Item    `gorm:"foreignKey:GroupId" json:"items"`
}
