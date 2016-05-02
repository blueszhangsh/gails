package gails

import (
	"log"

	"github.com/codegangsta/cli"
	"github.com/spf13/viper"
)

func Action(f cli.ActionFunc) cli.ActionFunc {
	return func(c *cli.Context) error {
		log.Println("Begin...")
		e := viper.ReadInConfig()
		if e == nil {
			e = f(c)
		}
		if e == nil {
			log.Println("Done!!!")
		} else {
			log.Println(e)
		}
		return e
	}
}

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
