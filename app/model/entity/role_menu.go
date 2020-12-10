package entity

type RoleMenu struct {
	RoleID int `gorm:"column:role_id;comment:角色id"`
	MenuID int `gorm:"column:menu_id;comment:菜单id"`
}

func (RoleMenu)TableName() string {
	return "role_menu"
}