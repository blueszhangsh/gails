package gails

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/BurntSushi/toml"
	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//Main entry
func Main(n string) error {
	app := cli.NewApp()
	app.Name = n
	app.Usage = fmt.Sprintf("Build by Gails web framework")
	app.Version = "v20160429"
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{

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
			Action: Action(func(*cli.Context) error {
				if IsProduction() {
					gin.SetMode(gin.ReleaseMode)
				}
				rt := gin.Default()
				for _, en := range engines {
					en.Mount(rt)
				}
				return rt.Run(fmt.Sprintf(":%d", viper.GetInt("http.port")))
			}),
		},
		{
			Name:    "routers",
			Aliases: []string{"ro"},
			Usage:   "print out all defined routes in match order, with names",
			Action: func(c *cli.Context) error {
				//TODO
				return nil
				// mux := martini.Classic()
				// web.Loop(func(en web.Engine) error {
				// 	en.Mount(mux)
				// 	return nil
				// })
				// for _, r := range mux.Router.All() {
				// 	fmt.Printf("%s\t%s\n", r.Method(), r.Pattern())
				// }
			},
		},

		{
			Name:    "database",
			Aliases: []string{"db"},
			Usage:   "database operations",
			Subcommands: []cli.Command{
				{
					Name:    "create",
					Aliases: []string{"n"},
					Usage:   "create database",
					Action: cfgAction(func(ctx *cli.Context) error {
						adp := viper.GetString("database.adapter")
						switch adp {
						case "postgres":
							c, a := psql("-c", fmt.Sprintf("CREATE DATABASE %s WITH ENCODING='UTF8'", viper.GetString("database.extras.dbname")))
							return Shell(c, a...)
						default:
							return fmt.Errorf("bad database adapter %s", adp)
						}
					}),
				},
				{
					Name:    "console",
					Aliases: []string{"c"},
					Usage:   "start a console for the database",
					Action: cfgAction(func(ctx *cli.Context) error {
						adp := viper.GetString("database.adapter")
						switch adp {
						case "postgres":
							c, a := psql(viper.GetString("database.extras.dbname"))
							return Shell(c, a...)
						default:
							return fmt.Errorf("bad database adapter %s", adp)
						}
					}),
				},
				{
					Name:    "drop",
					Aliases: []string{"d"},
					Usage:   "drop database",
					Action: cfgAction(func(ctx *cli.Context) error {
						adp := viper.GetString("database.adapter")
						switch adp {
						case "postgres":
							c, a := psql("-c", fmt.Sprintf("DROP DATABASE %s", viper.GetString("database.extras.dbname")))
							return Shell(c, a...)
						default:
							return fmt.Errorf("bad database adapter %s", adp)
						}
					}),
				},
				{
					Name:    "migrate",
					Aliases: []string{"m"},
					Usage:   "migrate the database",
					Action: Action(func(*cli.Context) error {
						for _, en := range engines {
							en.Migrate()
						}
						return nil
					}),
				},
				{
					Name:    "seed",
					Aliases: []string{"s"},
					Usage:   "load the seed data",
					Action: Action(func(*cli.Context) error {
						for _, en := range engines {
							en.Seed()
						}
						return nil
					}),
				},
			},
		},
		{
			Name:    "status",
			Aliases: []string{"st"},
			Usage:   "show status",
			Action: Action(func(*cli.Context) error {
				if IsProduction() {
					fmt.Println("=== CONFIG KEYS ===")
					fmt.Printf("%v\n", viper.AllKeys())

				} else {
					fmt.Println("=== CONFIG ITEMS ===")
					for k, v := range viper.AllSettings() {
						fmt.Printf("%s = %+v\n", k, v)
					}
				}

				fmt.Println("=== ENGINES ===")
				for _, en := range engines {
					vt := reflect.TypeOf(en).Elem()
					fmt.Printf("%s.%s\n", vt.PkgPath(), vt.Name())
				}
				return nil

			}),
		},
	}

	for _, en := range engines {
		cmds := en.Shell()
		app.Commands = append(app.Commands, cmds...)
	}

	return app.Run(os.Args)
}
