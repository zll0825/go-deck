package global

import "go-deck/pkg/zap"

var (
	Logger *zap.Logger
)

func InitAppLogger() {
	Logger = zap.NewZap(Config.LoggerConfig)
	Logger.Info("日志初始化成功")
}

