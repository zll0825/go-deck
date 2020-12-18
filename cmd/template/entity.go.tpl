package entity

import "time"

type {{.StructName}} struct {
    {{ range .Fields}}
        {{ if eq .FieldType "bool" }}
    {{.FieldName}}  *{{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}}{{- if .DataType -}};type:{{.DataType}}{{- if eq .FieldType "string" -}}{{- if .DataTypeLong -}}({{.DataTypeLong}}){{- end -}}{{- end -}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}{{- end -}}"`
        {{ else }}
    {{.FieldName}}  {{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}}{{- if .DataType -}};type:{{.DataType}}{{- if eq .FieldType "string" -}}{{- if .DataTypeLong -}}({{.DataTypeLong}}){{- end -}}{{- end -}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}{{- end -}}"`
        {{ end }}
    {{ end }}
}

{{ if .TableName }}
func ({{.StructName}}) TableName() string {
  return "{{.TableName}}"
}
{{ end }}