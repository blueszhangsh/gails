package base

import (
	"github.com/codegangsta/cli"
	"github.com/itpkg/gails"
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
	gails.Use(&Shell{})
}
