package gails

import (
	"log"

	"github.com/codegangsta/cli"
	"github.com/facebookgo/inject"
	"github.com/spf13/viper"
)

func Action(f cli.ActionFunc) cli.ActionFunc {
	return func(c *cli.Context) error {
		log.Println("Begin...")
		e := viper.ReadInConfig()
		var inj inject.Graph
		if e == nil {
			for _, en := range engines {
				if e = en.Map(&inj); e != nil {
					break
				}
			}
		}
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

	viper.SetDefault("http", map[string]interface{}{
		"port":   8080,
		"domain": "localhost",
		"ssl":    false,
	})
	viper.SetDefault("secrets", RandomStr(128))

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
}
