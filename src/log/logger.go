package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var L *zap.SugaredLogger

func init() {
	env := os.Getenv("ENV")
	var cfg zapcore.EncoderConfig
	f, err := os.OpenFile("/tmp/zap/log", os.O_RDWR|os.O_CREATE, os.ModeAppend)
	if err != nil {
		panic("failed to create log file")
	}
	if env == "DEV" {
		cfg = zap.NewDevelopmentEncoderConfig()
	} else {
		cfg = zap.NewProductionEncoderConfig()
	}
	fileEncoder := zapcore.NewJSONEncoder(cfg)

	core := zapcore.NewCore(fileEncoder, zapcore.AddSync(f), zap.DebugLevel)
	L = zap.New(core).Sugar()

}
