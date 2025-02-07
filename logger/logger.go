package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func New() *zap.Logger {
	stdout := zapcore.AddSync(os.Stdout)

	lvl := zap.NewAtomicLevelAt(zap.InfoLevel)

	devCfg := zap.NewDevelopmentEncoderConfig()
	devCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(devCfg)

	core := zapcore.NewCore(consoleEncoder, stdout, lvl)

	return zap.New(core)
}
