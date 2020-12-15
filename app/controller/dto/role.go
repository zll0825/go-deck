package dto

type CreateRole struct {
	RoleName string `json:"roleName" binding:"required"`
}

type UpdateRole struct {
	Id       int    `json:"id" binding:"required"`
	RoleName string `json:"roleName" binding:"required"`
}

type AddRoleMenu struct {
	RoleId  int   `json:"roleId" binding:"required"`
	MenuIds []int `json:"menuIds" binding:"required"`
}

type AddRoleApi struct {
	RoleId  int   `json:"roleId" binding:"required"`
	ApiIds []int `json:"apiIds" binding:"required"`
}
