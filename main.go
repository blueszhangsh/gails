package gails

import (
	"os"

	"github.com/codegangsta/cli"
)

//Main entry
func Main() error {
	app := cli.NewApp()
	app.Name = "itpkg"
	app.Usage = "IT-PACKAGE web framework"
	app.Version = "v20160401"
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
