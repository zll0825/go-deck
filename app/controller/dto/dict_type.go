package dto

type CreateDictType struct {
	Name        string `json:"name" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Description string `json:"description"`
	Status      int    `json:"status"`
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
