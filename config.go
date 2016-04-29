package gails

import "github.com/spf13/viper"

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
}
