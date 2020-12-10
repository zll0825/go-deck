package gorm

import (
	"bytes"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"io/ioutil"
	"sync"
)

var (
	mysqlList       []string
	gormManager     *Manager
	gormManagerOnce sync.Once
	openLog         sync.Map //name -> debug
)

// CheckValid 检验app.yaml中配置的mysql实例的配置正确性与连通性。
// 参数names是配置的实例的名称列表，如果为空，则检测所有配置的实例。
// 函数返回的error表示是否正常。
func CheckValid(names ...string) error {
	initManager()

	mysqlNames := names
	if len(mysqlNames) == 0 {
		mysqlNames = mysqlList
	}

	for _, name := range mysqlNames {
		if cli, err := getClient(name); err != nil {
			return err
		} else {
			// 获取通用数据库对象 sql.DB，然后使用其提供的功能
			sqlDB, err := cli.DB.DB()
			if err != nil {
				return err
			}
			if err := sqlDB.Ping(); err != nil {
				return err
			}
		}
	}
	return nil
}

// DBClient 返回某mysql实例对应的gorm官方DB对象。
// 参数names是mysql实例的名称。
// 备注: 服务初始化阶段调用CheckValid()检验过实例的有效性后，此方法访问此实例将不会再返回nil。
func DBClient(name string) *gorm.DB {
	cli, err := getClient(name)
	if err != nil {
		return nil
	}
	return cli.DB
}

func getClient(name string) (*Client, error) {
	initManager()

	client, err := gormManager.NewClient(name)
	if err != nil {
		return nil, err
	}

	//c.LogMode(false)
	//close gorm log info
	if v, ok := openLog.Load(name); ok {
		if debug, ok := v.(bool); ok && !debug {
			//client.SetLogger(log.New(ioutil.Discard, "", 0))
		}
	}

	return client, nil
}

type configWrap struct {
	GormConfigs map[string]*Config `yaml:"db"`
}
func initManager() {
	gormManagerOnce.Do(func() {
		cfgPath := "/Users/zll/Develop/go/src/github.com/zll0825/go-deck/conf/application.yaml"
		content, _ := ioutil.ReadFile(cfgPath)
		decoder := yaml.NewDecoder(bytes.NewReader(content))
		cfg := new(configWrap)
		_ = decoder.Decode(cfg)
		managerConfig := ManagerConfig(cfg.GormConfigs)

		gormManager = NewManager(&managerConfig)
	})
}

func fillDefaultConfig(conf *Config) {
	if conf.Driver == "" {
		conf.Driver = "mysql"
	}
	if conf.DialTimeout == 0 {
		conf.DialTimeout = 5000
	}
	if conf.ReadTimeout == 0 {
		conf.ReadTimeout = 5000
	}
	if conf.WriteTimeout == 0 {
		conf.WriteTimeout = 3000
	}
	if conf.MaxOpenConns == 0 {
		conf.MaxOpenConns = 256
	}
	if conf.MaxIdleConns == 0 {
		conf.MaxIdleConns = 10
	}
}
