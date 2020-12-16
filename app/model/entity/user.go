package entity

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primarykey"`
	Username  string    `json:"userName" gorm:"comment:用户登录名"`
	Password  string    `json:"-"  gorm:"comment:用户登录密码"`
	NickName  string    `json:"nickName" gorm:"comment:用户昵称" `
	HeaderImg string    `json:"headerImg" gorm:"comment:用户头像"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	Roles []Role `json:"roles" gorm:"many2many:user_role"`
}

func (User)TableName() string {
	return "user"
}
