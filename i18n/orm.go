package i18n

import (
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
	"golang.org/x/text/language"
)

//Locale locale model
type Locale struct {
	gorm.Model
	Lang    string `gorm:"not null;type:varchar(8);index"`
	Code    string `gorm:"not null;index;type:VARCHAR(255)"`
	Message string `gorm:"not null;type:varchar(800)"`
}

//OrmProvider db provider
type OrmProvider struct {
	Db     *gorm.DB        `inject:""`
	Logger *logging.Logger `inject:""`
}

//Set set locale
func (p *OrmProvider) Set(lng *language.Tag, code, message string) {
	var l Locale
	var err error
	if p.Db.Where("lang = ? AND code = ?", lng.String(), code).First(&l).RecordNotFound() {
		l.Lang = lng.String()
		l.Code = code
		l.Message = message
		err = p.Db.Create(&l).Error
	} else {
		l.Message = message
		err = p.Db.Save(&l).Error
	}
	if err != nil {
		p.Logger.Error(err)
	}
}

//Get get locale
func (p *OrmProvider) Get(lng *language.Tag, code string) string {
	var l Locale
	if err := p.Db.Where("lang = ? AND code = ?", lng.String(), code).First(&l).Error; err != nil {
		p.Logger.Error(err)
	}
	return l.Message

}

//Del del locale
func (p *OrmProvider) Del(lng *language.Tag, code string) {
	if err := p.Db.Where("lang = ? AND code = ?", lng.String(), code).Delete(Locale{}).Error; err != nil {
		p.Logger.Error(err)
	}
}

//Keys list locale keys
func (p *OrmProvider) Keys(lng *language.Tag) []string {
	var keys []string
	if err := p.Db.Model(&Locale{}).Where("lang = ?", lng.String()).Pluck("code", &keys).Error; err != nil {
		p.Logger.Error(err)
	}
	return keys
}
