package site

import (
	"github.com/codegangsta/cli"
	"github.com/spf13/viper"
)

func (p *Engine) Shell() []cli.Command {
	return []cli.Command{
	// {
	// 	Name:        "cache",
	// 	Aliases:     []string{"c"},
	// 	Usage:       "cache operations",
	// 	Subcommands: []cli.Command{
	// 	{
	// 		Name:    "list",
	// 		Aliases: []string{"l"},
	// 		Usage:   "list all cache items",
	// 		Flags:   []cli.Flag{config.ENV},
	// 		Action: config.InvokeAction(func(cp cache.Provider) error {
	// 			keys, err := cp.Status()
	// 			if err != nil {
	// 				return err
	// 			}
	// 			for k, v := range keys {
	// 				fmt.Printf("%s\t%d\n", k, v)
	// 			}
	// 			return nil
	// 		}),
	// 	},
	// 	{
	// 		Name:    "delete",
	// 		Aliases: []string{"d"},
	// 		Usage:   "delete item from cache",
	// 		Flags: []cli.Flag{
	// 			config.ENV,
	// 			cli.StringFlag{
	// 				Name:  "key, k",
	// 				Value: "",
	// 				Usage: "cache item's key",
	// 			},
	// 		},
	// 		Action: gails.Action(func(ctx *cli.Context) error {
	// 			k := ctx.String("key")
	// 			if k == "" {
	// 				return errors.New("key mustn't null")
	// 			}
	// 			_, err := mux.Invoke(func(cp cache.Provider) {
	// 				cp.Del(k)
	// 			})
	// 			return err
	// 		}),
	// 	},
	// 	{
	// 		Name:    "clear",
	// 		Aliases: []string{"c"},
	// 		Usage:   "delete all items from cache",
	// 		Action: config.InvokeAction(func(cp cache.Provider) error {
	// 			return cp.Clear()
	// 		}),
	// 	},
	// },
	// },
	}
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
