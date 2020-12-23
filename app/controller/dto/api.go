package dto

type CreateApi struct {
	Path        string `json:"path" binding:"required"`
	Description string `json:"description"`
	ApiGroup    string `json:"apiGroup" binding:"required"`
	Method      string `json:"method" binding:"required"`
}

type DeleteApi struct {
	DeleteReq
}

type UpdateApi struct {
	UpdateReq
	CreateApi
}

type SearchApi struct {
	PageReq
	CreateApi
}

type DetailApi struct {
	DetailReq
}
