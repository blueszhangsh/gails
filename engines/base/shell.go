package base

import (
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
	}
}

func init() {
	viper.SetDefault("database.username", "")	
	viper.SetDefault("redis.url", "redis://127.0.0.1:6379/0")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	gails.Use(&Shell{})
}
