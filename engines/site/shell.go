package site

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/BurntSushi/toml"
	"github.com/codegangsta/cli"
	"github.com/itpkg/gails"
	"github.com/spf13/viper"
)

func (p *Engine) Shell() []cli.Command {
	return []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "init config file",
			Action: func(*cli.Context) error {
				const fn = "config.toml"
				if _, err := os.Stat(fn); err == nil {
					msg := fmt.Sprintf("file %s already exists!", fn)
					log.Println(msg)
					return errors.New(msg)
				}

				args := viper.AllSettings()
				fd, err := os.Create(fn)
				if err != nil {
					log.Println(err)
					return err
				}
				defer fd.Close()
				end := toml.NewEncoder(fd)
				err = end.Encode(args)

				return err
			},
		},
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "start the web server",
			Action: gails.Action(func(*cli.Context) error {

				return nil
			}),
		},
		{
			Name:    "status",
			Aliases: []string{"st"},
			Usage:   "show status",
			Action: gails.Action(func(*cli.Context) error {
				if gails.IsProduction() {
					fmt.Println("=== CONFIG KEYS ===")
					fmt.Printf("%v\n", viper.AllKeys())

				} else {
					fmt.Println("=== CONFIG ITEMS ===")
					for k, v := range viper.AllSettings() {
						fmt.Printf("%s = %+v\n", k, v)
					}
				}

				fmt.Println("=== ENGINES ===")
				return gails.Each(func(en gails.Engine) error {
					vt := reflect.TypeOf(en).Elem()
					fmt.Printf("%s.%s\n", vt.PkgPath(), vt.Name())
					return nil
				})
			}),
		},
	}
}

func init() {
	viper.SetDefault("http", map[string]interface{}{
		"port": 8080,
	})
	viper.SetDefault("secrets", gails.RandomStr(128))
	viper.SetDefault(
		"database",
		map[string]interface{}{
			"adapter": "postgres",
			"pool": map[string]int{
				"max_open": 100,
				"max_idle": 5,
			},
			"extras": map[string]interface{}{
				"host":     "localhost",
				"user":     "postgres",
				"password": "",
				"dbname":   "gails_dev",
				"sslmode":  "disable",
			},
		},
	)

	viper.SetDefault(
		"redis",
		map[string]interface{}{
			"host": "localhost",
			"port": 6379,
			"db":   0,
		})
}
