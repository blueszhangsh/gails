package base

import (
	"github.com/itpkg/gails"
	"github.com/itpkg/gails/i18n"
	"github.com/itpkg/gails/setting"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
)

type Migration struct {
	Db     *gorm.DB        `inject:""`
	Logger *logging.Logger `inject:""`
}

func (p *Migration) Migrate() {
	db := p.Db
	db.AutoMigrate(
		&setting.Model{}, &i18n.Locale{},
		&Notice{},
	)
	db.Model(&i18n.Locale{}).AddUniqueIndex("idx_locales_lang_code", "lang", "code")
}

func (p *Migration) Seed() {

}

func init() {
	gails.Use(&Migration{})
}
