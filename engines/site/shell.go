package site

import (
	"fmt"
	"reflect"

	"github.com/codegangsta/cli"
	"github.com/itpkg/gails"
	"github.com/spf13/viper"
)

//Shell base shell
type Shell struct {
}

//Commands base commands
func (p *Shell) Commands() []cli.Command {
	return []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "init config file",
		},
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "start the web server",
		},
		{
			Name:    "status",
			Aliases: []string{"sts"},
			Usage:   "show status",
			Action: func(*cli.Context) {
				if gails.IsProduction() {
					fmt.Println("=== CONFIG KEYS ===")
					for _, v := range viper.AllKeys() {
						fmt.Println(v)
					}
				} else {
					fmt.Println("=== CONFIG ITEMS ===")
					for k, v := range viper.AllSettings() {
						fmt.Printf("%s = %+v\n", k, v)
					}
				}

				fmt.Println("=== BEANS ===")
				gails.Loop(func(n string, o interface{}) error {
					vt := reflect.TypeOf(o).Elem()
					fmt.Printf("name = %s, type = %s.%s\n", n, vt.PkgPath(), vt.Name())
					return nil
				})
			},
		},
	}
}

func init() {
	viper.SetDefault("database.username", "")
	viper.SetDefault("redis.url", "redis://127.0.0.1:6379/0")

	gails.Use(&Shell{})
}
