package auth

import (
	"github.com/itpkg/gails"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
)

type Dao struct {
	Logger    *logging.Logger  `inject:""`
	Db        *gorm.DB         `inject:""`
	Encryptor *gails.Encryptor `inject:""`
}
