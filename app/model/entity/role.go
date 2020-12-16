package entity

import (
	"time"
)

type Role struct {
	ID   int    `gorm:"primarykey"`
	Name string `json:"name" gorm:"comment:角色名"`
	Key  string `json:"key" gorm:"comment:角色key"`

	Apis  []Api  `json:"apis" gorm:"many2many:role_id"`
	Menus []Menu `json:"menus" gorm:"many2many:role_id"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (Role) TableName() string {
	return "role"
}
