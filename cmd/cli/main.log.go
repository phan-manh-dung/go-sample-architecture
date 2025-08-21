package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {

	// 1.
	sugar := zap.NewExample().Sugar()
	sugar.Infof("Hello name:%s, age:%d", "Dũng", 20)

	// logger
	logger := zap.NewExample()
	logger.Info("Hello", zap.String("name", "Dũng"), zap.Int("age", 20))

	//2. các config
	logger2 := zap.NewExample()
	logger2.Info("Hello New Example")

	// Development
	logger2, _ = zap.NewDevelopment()
	logger2.Info("Hello New Development") // 2025-08-21T11:06:42.929+0700    INFO    cli/main.log.go:19      Hello New Development

	// Production
	logger2, _ = zap.NewProduction()
	logger2.Info("Hello New Production") // {"level":"info","ts":1755749202.9467287,"caller":"cli/main.log.go:23","msg":"Hello New Production"}

	//3. Sử dụng 1 hàm ghi log vào 1 file
	encoder := getEncoderLog()
	writer := getWriterSynce()
	core := zapcore.NewCore(encoder, writer, zapcore.InfoLevel)
	logger3 := zap.New(core)
	logger3.Info("Info log", zap.Int("line", 1))
	logger3.Error("Info log", zap.Int("line", 2))
}

// format log
func getEncoderLog() zapcore.Encoder { // zapcore là tính năng đóng gói của zap , Encoder là định dạng
	encoderConfig := zap.NewProductionEncoderConfig()

	// 1755749202.9467287 -> 2025-08-21T11:06:42.929+0700
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// ts -> Time
	encoderConfig.TimeKey = "time"
	// from info INFO
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// "caller":"cli/main.log.go:23"
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriterSynce() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)
}
