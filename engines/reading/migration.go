package reading

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
	db := p.Db
	db.AutoMigrate(&Note{})
	db.Model(&Note{}).AddUniqueIndex("idx_reading_notes_user_title", "user_id", "title")
}

func (p *Migration) Seed() {

}

func init() {
	gails.Use(&Migration{})
}
