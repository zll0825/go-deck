package dto

type PageReq struct {
	Page int `json:"page" form:"page" binding:"required"`
	Size int `json:"size" form:"size" binding:"required"`
}

type DeleteReq struct {
	Ids []int `json:"ids" form:"ids" binding:"required"`
}

type UpdateReq struct {
	Id int `json:"id" form:"id" binding:"required"`
}

type DetailReq struct {
	Id int `json:"id" form:"id" binding:"required"`
}
