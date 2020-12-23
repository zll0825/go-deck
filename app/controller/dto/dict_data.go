package dto

type CreateDictData struct {
	TypeId      int    `json:"typeId" binding:"required"`
	Label       string `json:"label" binding:"required"`
	Value       string `json:"value" binding:"required"`
	Description string `json:"description"`
	Status      int    `json:"status"`
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
