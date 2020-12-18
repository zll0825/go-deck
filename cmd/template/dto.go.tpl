package dto

type Create{{.StructName}} struct {

}

type Delete{{.StructName}} struct {
	DeleteReq
}

type Search{{.StructName}} struct {
	PageReq
	Create{{.StructName}}
}

type Detail{{.StructName}} struct {
	DetailReq
}

type Update{{.StructName}} struct {
	UpdateReq
	Create{{.StructName}}
}