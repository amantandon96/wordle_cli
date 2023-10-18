package log

import (
	"go.uber.org/zap"
	"wordle_cli/config"
	"wordle_cli/flags"
)

var L *zap.SugaredLogger

func init() {
	cfg := zap.NewProductionConfig()
	if config.V.GetBool(flags.Debug) {
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}
	log, err := cfg.Build()
	if err != nil {
		panic("failed to init logger")
	}
	L = log.Sugar()

}
