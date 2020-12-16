package zap

import (
	"fmt"
	rotate "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

var level zapcore.Level

type Config struct {
	Level         string `yaml:"level"`
	Format        string `yaml:"format"`
	Prefix        string `yaml:"prefix"`
	Directory     string `yaml:"directory"`
	LinkName      string `yaml:"linkName"`
	ShowLine      bool   `yaml:"showLine"`
	EncodeLevel   string `yaml:"encodeLevel"`
	StacktraceKey string `yaml:"stacktraceKey"`
	LogInConsole  bool   `yaml:"logInConsole"`
}

type Logger struct {
	*zap.Logger

	config *Config
}

func NewZap(config *Config) *Logger {
	_, err := os.Stat(config.Directory)
	if err != nil { // 判断是否有Directory文件夹
		fmt.Printf("create %v directory\n", config.Directory)
		_ = os.Mkdir(config.Directory, os.ModePerm)
	}

	switch config.Level { // 初始化配置文件的Level
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	logger := &Logger{
		nil,
		config,
	}
	if level == zap.DebugLevel || level == zap.ErrorLevel {
		logger.Logger = zap.New(logger.getEncoderCore(), zap.AddStacktrace(level))
	} else {
		logger.Logger = zap.New(logger.getEncoderCore())
	}
	if config.ShowLine {
		logger.Logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

// getEncoderConfig 获取zapcore.EncoderConfig
func (l *Logger) getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  config.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     l.CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case l.config.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case l.config.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case l.config.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case l.config.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// getEncoder 获取zapcore.Encoder
func (l *Logger) getEncoder() zapcore.Encoder {
	if l.config.Format == "json" {
		return zapcore.NewJSONEncoder(l.getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(l.getEncoderConfig())
}

// getEncoderCore 获取Encoder的zapcore.Core
func (l *Logger) getEncoderCore() (core zapcore.Core) {
	writer, err := l.getWriteSyncer() // 使用file-rotatelogs进行日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return
	}
	return zapcore.NewCore(l.getEncoder(), writer, level)
}

// 自定义日志输出时间格式
func (l *Logger) CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(l.config.Prefix + "2006/01/02 15:04:05.000"))
}

func (l *Logger) getWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := rotate.New(
		path.Join(l.config.Directory, "%Y-%m-%d.log"),
		rotate.WithLinkName(l.config.LinkName),
		rotate.WithMaxAge(7*24*time.Hour),
		rotate.WithRotationTime(24*time.Hour),
	)
	if l.config.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
