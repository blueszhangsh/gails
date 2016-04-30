package hr

import (
	"github.com/itpkg/gails"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
)

type Migration struct {
	Db     *gorm.DB        `inject:""`
	Logger *logging.Logger `inject:""`
}

func (p *Migration) Migrate() {

}

func (p *Migration) Seed() {

}

func init() {
	gails.Use(&Migration{})
}
