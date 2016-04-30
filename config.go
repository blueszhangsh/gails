package gails

import "github.com/spf13/viper"

func IsProduction() bool {
	return viper.GetString("env") == "production"
}

func init() {
	viper.SetEnvPrefix("gails")
	viper.BindEnv("env")
	viper.SetDefault("env", "development")

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
}
