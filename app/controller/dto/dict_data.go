package dto

type CreateDictData struct {
	TypeID int    `gorm:""`
	Label  string `json:"name" form:"name" gorm:"column:name;comment:字典名（中）"`
	Value  string `json:"type" form:"type" gorm:"column:type;comment:字典名（英）"`
	Desc   string `json:"desc" form:"desc" gorm:"column:desc;comment:描述"`
	Status int    `json:"status" form:"status" gorm:"column:status;comment:状态"`
}

type DeleteDictData struct {
	DeleteReq
}

type UpdateDictData struct {
	UpdateReq
	CreateDictData
}

type SearchDictData struct {
	PageReq
	CreateDictData
}

type DetailDictData struct {
	DetailReq
}
