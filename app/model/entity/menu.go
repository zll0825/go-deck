package entity

import (
	"time"
)

type Menu struct {
	ID          int    `gorm:"primarykey"`
	MenuLevel   int    `json:"-"`
	ParentId    int    `json:"parentId" gorm:"comment:父菜单ID"`
	Path        string `json:"path" gorm:"comment:路由path"`
	Name        string `json:"name" gorm:"comment:路由name"`
	Hidden      bool   `json:"hidden" gorm:"comment:是否在列表隐藏"`
	Component   string `json:"component" gorm:"comment:对应前端文件路径"`
	Sort        int    `json:"sort" gorm:"comment:排序标记"`
	KeepAlive   bool   `json:"keepAlive" gorm:"comment:是否缓存"`
	DefaultMenu bool   `json:"defaultMenu" gorm:"comment:是否是基础路由（开发中）"`
	Title       string `json:"title" gorm:"comment:菜单名"`
	Icon        string `json:"icon" gorm:"comment:菜单图标"`

	Children []Menu `json:"children" gorm:"-"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (Menu) TableName() string {
	return "menu"
}
