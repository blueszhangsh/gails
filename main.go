package gails

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

//Main entry
func Main(n string) error {
	app := cli.NewApp()
	app.Name = n
	app.Usage = fmt.Sprintf("Build by Gails web framework")
	app.Version = "v20160429"
	app.EnableBashCompletion = true
	app.Commands = make([]cli.Command, 0)

	for _, en := range engines {
		cmds := en.Shell()
		app.Commands = append(app.Commands, cmds...)
	}

	return app.Run(os.Args)
}
