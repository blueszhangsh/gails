package gails

import (
	"github.com/codegangsta/cli"
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
)

type Engine interface {
	Map(*inject.Graph) error
	Mount(*gin.Engine)
	Assets() error
	Migrate()
	Seed()
	Shell() []cli.Command
}

var engines []Engine

//Register register engine
func Register(ens ...Engine) {
	engines = append(engines, ens...)
}
