package gails

import (
	"github.com/codegangsta/cli"
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
)

type Engine interface {
	Map(*inject.Graph)
	Mount(*gin.Engine)
	Assets() error
	Migrate()
	Seed()
	Shell() []cli.Command
}

var engines []Engine

func Register(ens ...Engine) {
	engines = append(engines, ens...)
}

func Each(f func(Engine) error) error {
	for _, en := range engines {
		if er := f(en); er != nil {
			return er
		}
	}
	return nil
}
