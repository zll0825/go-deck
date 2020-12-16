package dto

type CreateRole struct {
	Name string `json:"name" binding:"required"`
	Key string `json:"key" binding:"required"`
}

type UpdateRole struct {
	Id       int    `json:"id" binding:"required"`
	RoleName string `json:"name" binding:"required"`
}

type BindRoleMenu struct {
	RoleId  int   `json:"roleId" binding:"required"`
	MenuIds []int `json:"menuIds" binding:"required"`
}

type BindRoleApi struct {
	RoleId  int   `json:"roleId" binding:"required"`
	ApiIds []int `json:"apiIds" binding:"required"`
}
