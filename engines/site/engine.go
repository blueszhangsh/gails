package site

import (
	"github.com/facebookgo/inject"
	"github.com/itpkg/gails"
	"github.com/itpkg/gails/setting"
	"github.com/itpkg/web/i18n"
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
	db := p.Db
	db.AutoMigrate(
		&setting.Model{}, &i18n.Locale{},
		&Notice{},
	)
	db.Model(&i18n.Locale{}).AddUniqueIndex("idx_locales_lang_code", "lang", "code")

}
func (p *Engine) Seed() {

}

func init() {
	gails.Register(&Engine{})
}
