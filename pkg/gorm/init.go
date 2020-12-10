package gorm

import (
	// for mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// TODO: We should use mysql cluster proxy instead of local DefaultMgr!!!
var (
	DefaultMgr *Manager
)

func init() {
	DefaultMgr = NewManager(nil)
}
