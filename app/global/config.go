package global

import (
	"bytes"
	"go-deck/pkg/casbin"
	"go-deck/pkg/gorm"
	"go-deck/pkg/jwt"
	"go-deck/pkg/zap"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	Config AppConfig
)

type AppConfig struct {
	LoggerConfig *zap.Config             `yaml:"logger"`
	DBConfig     map[string]*gorm.Config `yaml:"db"`
	JwtConfig    *jwt.Config             `yaml:"jwt"`
	CasbinConfig *casbin.Config          `yaml:"casbin"`
	SysConfig    *SysConfig              `yaml:"system"`
}

type SysConfig struct {
	Env  string `yaml:"env"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func InitAppConfig(cfgPath string) {
	content, err := ioutil.ReadFile(cfgPath)

	if err != nil {
		panic(err)
	}

	decoder := yaml.NewDecoder(bytes.NewReader(content))
	err = decoder.Decode(&Config)
	if err != nil {
		panic(err)
	}
}
