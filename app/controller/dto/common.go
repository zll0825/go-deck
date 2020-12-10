package dto

type PageReq struct {
	Page int `json:"page" form:"page"`
	Size int `json:"size" form:"size"`
}

type DetailReq struct {
	Id int `json:"id" form:"id"`
}

type UpdateReq struct {
	Id int `json:"id" form:"id"`
}

type DeleteReq struct {
	Ids []int `json:"ids" form:"ids"`
}
