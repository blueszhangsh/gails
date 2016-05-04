package team

import (
	"github.com/codegangsta/cli"
	"github.com/facebookgo/inject"
	"github.com/itpkg/gails"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
)

type Engine struct {
	Db     *gorm.DB        `inject:""`
	Logger *logging.Logger `inject:""`
}

func (p *Engine) Map(*inject.Graph) {

}
func (p *Engine) Assets() error {
	return nil
}
func (p *Engine) Migrate() {

}
func (p *Engine) Seed() {

}
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{}
}

func init() {
	gails.Register(&Engine{})
}
