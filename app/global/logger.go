package global

import "go-deck/pkg/zap"

var (
	Logger *zap.Logger
)

func InitAppLogger() {
	Logger = zap.NewZap(Config.LoggerConfig)
}

