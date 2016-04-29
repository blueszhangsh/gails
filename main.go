package gails

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/spf13/viper"
)

//Main entry
func Main(n string) error {
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	app := cli.NewApp()
	app.Name = n
	app.Usage = fmt.Sprintf("Build by Gails web framework")
	app.Version = "v20160429"
	app.Commands = make([]cli.Command, 0)
	Loop(func(n string, o interface{}) error {
		switch o := o.(type) {
		case Shell:
			cmds := o.Commands()
			app.Commands = append(app.Commands, cmds...)
		}
		return nil
	})

	return app.Run(os.Args)
}
