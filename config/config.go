package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var V *viper.Viper

func init() {
	V = viper.NewWithOptions()
	V.SetConfigFile(os.Getenv("CONFIG_PATH"))
	err := V.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err.Error()))
	}
}
