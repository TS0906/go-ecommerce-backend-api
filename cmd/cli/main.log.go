package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello, name: %s, age: %d", "thinh", 30) // like fmt.Printf

	// //logger
	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "thinh"), zap.Int("age", 30)) // key-value

	// logger := zap.NewExample() // in ra log co ban
	// logger.Info("Hello")

	// //Development in ra log chi tiết
	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello Development")

	// //production in ra log ngắn gọn
	// logger, _ = zap.NewProduction()
	// logger.Info("Hello Production")

	//encode
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))
}

func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder // thoi gian tu encode thanh dang iso8601
	// ts -> time stamp
	encodeConfig.TimeKey = "time" // key cua thoi gian

	//from info INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder // in ra INFO

	//caller":"cli/main.log.go:21
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder // in ra duong dan file
	return zapcore.NewJSONEncoder(encodeConfig)
}

func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)
}
