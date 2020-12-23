package dto

type CreateMenu struct {
	MenuLevel   int    `json:"-"`
	ParentId    int    `json:"parentId" binding:"required"`
	Path        string `json:"path" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Hidden      bool   `json:"hidden" binding:"required"`
	Component   string `json:"component" binding:"required"`
	Sort        int    `json:"sort"`
	KeepAlive   bool   `json:"keepAlive"`
	DefaultMenu bool   `json:"defaultMenu"`
	Title       string `json:"title"`
	Icon        string `json:"icon"`
}

type DeleteMenu struct {
	DeleteReq
}

type UpdateMenu struct {
	UpdateReq
	CreateMenu
}

type SearchMenu struct {
	PageReq
	CreateMenu
}

type DetailMenu struct {
	DetailReq
}
