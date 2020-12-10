package global

import (
	"go-deck/app/model"
)

var (
	DB *model.Dao
)

func InitDB()  {
	DB = model.NewDao()
}