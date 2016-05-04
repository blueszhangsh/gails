package reading

import (
	"github.com/facebookgo/inject"
	"github.com/itpkg/gails"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
)

type Engine struct {
	Db     *gorm.DB        `inject:""`
	Logger *logging.Logger `inject:""`
}

func (p *Engine) Map(*inject.Graph) error {
	return nil

}
func (p *Engine) Assets() error {
	return nil
}
func (p *Engine) Migrate() {
	db := p.Db
	db.AutoMigrate(&Note{})
	db.Model(&Note{}).AddUniqueIndex("idx_reading_notes_user_title", "user_id", "title")

}
func (p *Engine) Seed() {

}

func init() {
	gails.Register(&Engine{})
}
