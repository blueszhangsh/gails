package site

import (
	"github.com/codegangsta/cli"
	"github.com/spf13/viper"
)

func (p *Engine) Shell() []cli.Command {
	return []cli.Command{}
}

func init() {
	viper.SetDefault(
		"database",
		map[string]interface{}{
			"adapter": "postgres",
			"pool": map[string]int{
				"max_open": 100,
				"max_idle": 5,
			},
			"extras": map[string]interface{}{
				"host":    "localhost",
				"user":    "postgres",
				"dbname":  "gails_dev",
				"sslmode": "disable",
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
