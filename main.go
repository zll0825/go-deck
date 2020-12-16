package main

import (
	"flag"
	"fmt"
	"go-deck/app/global"
	"go-deck/app/router"
	"os"
	"path"
	"path/filepath"
)

var (
	runMode string
	cfgPath string
)

func main() {
	// 接收启动参数
	flag.StringVar(&runMode, "runMode", "dev", "-runMode=[dev|qa|prd]")
	flag.StringVar(&cfgPath, "cfgPath", "", "-cfgPath=/path/to/[config/application.yml]")
	flag.Parse()

	fmt.Printf("starting runMode:%v, cfgPath:%v", runMode, cfgPath)

	// 初始化配置文件
	if cfgPath == "" {
		root, _ := os.Getwd()

		cfgPath = filepath.Join(root, "conf/application.yaml")
	} else {
		cfgPath = path.Clean(cfgPath)
	}
	global.InitAppConfig(cfgPath)
	// 初始化日志
	global.InitAppLogger()
	// 初始化数据库连接
	global.InitDB()
	// 初始化路由
	server := router.Routers()
	_ = server.Run(fmt.Sprintf("%s:%d", global.Config.SysConfig.Host, global.Config.SysConfig.Port))

	fmt.Printf("gateway running")
}
