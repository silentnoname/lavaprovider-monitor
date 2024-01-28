package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var Log *zap.Logger

// InitLog  initializes the logger
func InitLog() {
	fileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./alert.log",
		MaxSize:    100,
		MaxBackups: 3,
		MaxAge:     30,
		Compress:   false,
	})
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	fileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(fileWriteSyncer, zapcore.AddSync(os.Stdout)), zapcore.DebugLevel)
	Log = zap.New(fileCore, zap.AddCaller())
}
