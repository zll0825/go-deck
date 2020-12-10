package casbin

import (
	"gorm.io/gorm"
	"os"
	"path/filepath"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/gorm-adapter/v3"
)

type Casbin struct {
	config *Config
	DB *gorm.DB
}

type Config struct {
	Path string `yaml:"path"`
}

func NewCasbin(config *Config, db *gorm.DB) (e *casbin.Enforcer, err error) {
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		return
	}

	if config.Path == "" {
		root, _ := os.Getwd()
		config.Path = filepath.Join(root + "/conf/rbac_model.conf")
	}
	e, err = casbin.NewEnforcer(config.Path, adapter)
	if err != nil {
		return
	}

	if err = e.LoadPolicy(); err != nil {
		return
	}
	return
}
