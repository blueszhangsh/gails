package gails

import "github.com/jinzhu/gorm"

//Migration migration
type Migration interface {
	Up(*gorm.DB)
	Down(*gorm.DB)
	Seed(*gorm.DB)
}
