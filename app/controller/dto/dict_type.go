package dto

type CreateDictType struct {
	Name   string `json:"name" form:"name" gorm:"column:name;comment:字典名（中）"`
	Type   string `json:"type" form:"type" gorm:"column:type;comment:字典名（英）"`
	Desc   string `json:"desc" form:"desc" gorm:"column:desc;comment:描述"`
	Status int    `json:"status" form:"status" gorm:"column:status;comment:状态"`
}

type DeleteDictType struct {
	DeleteReq
}

type UpdateDictType struct {
	UpdateReq
	CreateDictType
}

type SearchDictType struct {
	PageReq
	CreateDictType
}

type DetailDictType struct {
	DetailReq
}
