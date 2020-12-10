package model

import (
	myGorm "go-deck/pkg/gorm"
	"gorm.io/gorm"
)

type Dao struct {
	System     *gorm.DB
	BusinessDB *gorm.DB
}

func NewDao() *Dao {
	return &Dao{
		System:     myGorm.DBClient("system"),
		BusinessDB: myGorm.DBClient("business_db"),
	}
}

func SystemDB() *gorm.DB {
	return myGorm.DBClient("system")
}
