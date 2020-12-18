package cmd

import (
	"fmt"
	"go-deck/app/global"
	"go-deck/app/router"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	runMode string
	cfgPath string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		// 接收启动参数
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
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&runMode, "runMode", "m", "dev", "-runMode=[dev|qa|prd]")
	rootCmd.PersistentFlags().StringVarP(&cfgPath, "cfgPath", "c", "", "-cfgPath=/path/to/[config/application.yml]")
}
