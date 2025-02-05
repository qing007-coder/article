package logger

import (
	"article/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"time"
)

func LoggerInit(conf *config.GlobalConfig) *zap.SugaredLogger {
	logMode := zapcore.DebugLevel

	core := zapcore.NewCore(getEncoder(), getWriteSyncer(conf), logMode)
	return zap.New(core).Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Local().Format(time.DateTime))
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriteSyncer(conf *config.GlobalConfig) zapcore.WriteSyncer {
	stSeparator := string(filepath.Separator)
	stRootDir, _ := os.Getwd()
	stLogFilePath := stRootDir + stSeparator + "Logger" + stSeparator + "log" + stSeparator + time.Now().Format(time.DateOnly) + ".txt"
	//fmt.Println(stLogFilePath)

	lumberjackSyncer := lumberjack.Logger{
		Filename:   stLogFilePath,
		MaxSize:    conf.Logger.MaxSize,
		MaxBackups: conf.Logger.MaxBackups,
		MaxAge:     conf.Logger.MaxAge,
		Compress:   true,
	}
	return zapcore.AddSync(&lumberjackSyncer)
}
