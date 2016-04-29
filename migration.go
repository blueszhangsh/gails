package gails

import "github.com/jinzhu/gorm"

//Migration migration
type Migration interface {
	Migrate(*gorm.DB)
	Seed(*gorm.DB)
}
