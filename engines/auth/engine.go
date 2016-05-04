package auth

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
	db := p.Db
	db.AutoMigrate(
		&User{}, &Role{}, &Permission{}, &Log{},
	)
	db.Model(&User{}).AddUniqueIndex("idx_user_provider_type_id", "provider_type", "provider_id")
	db.Model(&Role{}).AddUniqueIndex("idx_roles_name_resource_type_id", "name", "resource_type", "resource_id")
	db.Model(&Permission{}).AddUniqueIndex("idx_permissions_user_role", "user_id", "role_id")

}
func (p *Engine) Seed() {

}
func (p *Engine) Shell() []cli.Command {
	return []cli.Command{}
}

func init() {
	gails.Register(&Engine{})
}
