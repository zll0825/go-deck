package entity

import (
	"time"
)

type Role struct {
	ID     int    `gorm:"primarykey"`
	Name   string `json:"name" gorm:"comment:角色名"`
	UserID int    `json:"userId" gorm:"comment:用户ID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (Role)TableName() string {
	return "role"
}