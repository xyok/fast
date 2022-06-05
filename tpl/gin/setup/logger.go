package setup

import (
	"{{ .AppName }}/conf"
	"io"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// setupLogger 初始化日志配置
func setupLogger(level string, writer io.Writer, encoder zapcore.Encoder) (err error) {
	var loggerLevel = new(zapcore.Level)
	err = loggerLevel.UnmarshalText([]byte(level))
	if err != nil {
		return
	}
	core := zapcore.NewCore(encoder, zapcore.AddSync(writer), loggerLevel)
	// 替换zap包中全局的logger实例，后续直接调 global.Logger
	zap.ReplaceGlobals(zap.New(core, zap.AddCaller()))
	return err
}

// DebugLogger 日志设置为控制台标准输出
func DebugLogger(cfg conf.LoggerCfg) (err error) {
	return setupLogger(
		cfg.Level,
		os.Stdout,
		zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
	)
}

// ReleaseLogger 日志格式化为json并输出到日志
func ReleaseLogger(cfg conf.LoggerCfg) (err error) {
	return setupLogger(
		cfg.Level,
		&lumberjack.Logger{
			Filename:   cfg.Filepath,
			MaxSize:    cfg.MaxSize,
			MaxBackups: cfg.MaxBackup,
			MaxAge:     cfg.MaxAge,
		},
		zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}),
	)
}
