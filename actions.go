package gails

import (
	"log"

	"github.com/codegangsta/cli"
	"github.com/facebookgo/inject"
	"github.com/spf13/viper"
)

func cfgAction(f cli.ActionFunc) cli.ActionFunc {
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

func Action(f cli.ActionFunc) cli.ActionFunc {
	return cfgAction(func(ctx *cli.Context) error {

		logger := Logger()
		var inj inject.Graph
		if !IsProduction() {
			inj.Logger = logger
		}

		inj.Provide(
			&inject.Object{
				Value: logger,
			},
		)
		for _, en := range engines {
			if er := en.Map(&inj); er != nil {
				return er
			}
			inj.Provide(&inject.Object{Value: en})
		}
		if err := inj.Populate(); err != nil {
			return err
		}

		return f(ctx)

	})
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
