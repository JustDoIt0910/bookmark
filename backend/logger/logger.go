package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var SugarLogger *zap.SugaredLogger

func InitLogger()  {
	encoder := getEncoder()
	writer := getLogWriter()
	core := zapcore.NewCore(encoder, writer, zap.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	SugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberjackWriter := &lumberjack.Logger{
		Filename: "./log/log.log",
		MaxSize: 10,
		MaxAge: 30,
		MaxBackups: 5,
		Compress: false,
	}
	return zapcore.AddSync(lumberjackWriter)
}