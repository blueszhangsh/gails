package gails

import "github.com/codegangsta/cli"

//Shell shell commands
type Shell interface {
	Commands() []cli.Command
}
