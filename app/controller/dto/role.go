package dto

type CreateRole struct {
	Name string `json:"name" binding:"required"`
	Key string `json:"key" binding:"required"`
}

type DeleteRole struct {
	DeleteReq
}

type UpdateRole struct {
	UpdateReq
	CreateRole
}

type SearchRole struct {
	PageReq
	CreateRole
}

type DetailRole struct {
	DetailReq
}

type BindRoleMenu struct {
	RoleId  int   `json:"roleId" binding:"required"`
	MenuIds []int `json:"menuIds" binding:"required"`
}

type BindRoleApi struct {
	RoleId  int   `json:"roleId" binding:"required"`
	ApiIds []int `json:"apiIds" binding:"required"`
}
