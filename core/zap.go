package core

import (
	"os"
	"server-fiber/core/internal"
	"server-fiber/global"
	"server-fiber/utils"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Zap 获取 zap.Logger
// Author wuhao
func zapInit() (logger *zap.Logger, err error) {
	ok, err := utils.PathExists(global.CONFIG.Zap.Director)
	if !ok { // 判断是否有Director文件夹
		// log.Printf("create %v directory\n", global.CONFIG.Zap.Director)
		err = os.Mkdir(global.CONFIG.Zap.Director, os.ModePerm)
		return
	}

	if err != nil {
		return
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return
}

// func Zap() (logger *zap.Logger) {
// 	if ok, _ := utils.PathExists(global.CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
// 		fmt.Printf("create %v directory\n", global.CONFIG.Zap.Director)
// 		_ = os.Mkdir(global.CONFIG.Zap.Director, os.ModePerm)
// 	}
// 	// 调试级别
// 	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
// 		return lev == zap.DebugLevel
// 	})
// 	// 日志级别
// 	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
// 		return lev == zap.InfoLevel
// 	})
// 	// 警告级别
// 	warnPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
// 		return lev == zap.WarnLevel
// 	})
// 	// 错误级别
// 	errorPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
// 		return lev >= zap.ErrorLevel
// 	})

// 	var path = global.CONFIG.Zap.Director

// 	cores := [...]zapcore.Core{
// 		getEncoderCore(fmt.Sprintf("./%s/server_debug.log", path), debugPriority),
// 		getEncoderCore(fmt.Sprintf("./%s/server_info.log", path), infoPriority),
// 		getEncoderCore(fmt.Sprintf("./%s/server_warn.log", path), warnPriority),
// 		getEncoderCore(fmt.Sprintf("./%s/server_error.log", path), errorPriority),
// 	}
// 	logger = zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())

// 	if global.CONFIG.Zap.ShowLine {
// 		logger = logger.WithOptions(zap.AddCaller())
// 	}
// 	return logger
// }

// getEncoderConfig 获取zapcore.EncoderConfig
// func getEncoderConfig() (config zapcore.EncoderConfig) {
// 	config = zapcore.EncoderConfig{
// 		MessageKey:     "message",
// 		LevelKey:       "level",
// 		TimeKey:        "time",
// 		NameKey:        "logger",
// 		CallerKey:      "caller",
// 		StacktraceKey:  global.CONFIG.Zap.StacktraceKey,
// 		LineEnding:     zapcore.DefaultLineEnding,
// 		EncodeLevel:    zapcore.LowercaseLevelEncoder,
// 		EncodeTime:     CustomTimeEncoder,
// 		EncodeDuration: zapcore.SecondsDurationEncoder,
// 		EncodeCaller:   zapcore.FullCallerEncoder,
// 	}
// 	switch {
// 	case global.CONFIG.Zap.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
// 		config.EncodeLevel = zapcore.LowercaseLevelEncoder
// 	case global.CONFIG.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
// 		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
// 	case global.CONFIG.Zap.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
// 		config.EncodeLevel = zapcore.CapitalLevelEncoder
// 	case global.CONFIG.Zap.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
// 		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
// 	default:
// 		config.EncodeLevel = zapcore.LowercaseLevelEncoder
// 	}
// 	return config
// }

// getEncoder 获取zapcore.Encoder
// func getEncoder() zapcore.Encoder {
// 	if global.CONFIG.Zap.Format == "json" {
// 		return zapcore.NewJSONEncoder(getEncoderConfig())
// 	}
// 	return zapcore.NewConsoleEncoder(getEncoderConfig())
// }

// getEncoderCore 获取Encoder的zapcore.Core
// func getEncoderCore(fileName string, level zapcore.LevelEnabler) (core zapcore.Core) {
// 	writer := utils.GetWriteSyncer(fileName) // 使用file-rotatelogs进行日志分割
// 	return zapcore.NewCore(getEncoder(), writer, level)
// }

// 自定义日志输出时间格式
// func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
// 	enc.AppendString(t.Format(global.CONFIG.Zap.Prefix + "2006/01/02 - 15:04:05.000"))
// }
