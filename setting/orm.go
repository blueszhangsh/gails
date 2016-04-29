package setting

import (
	"github.com/itpkg/gails"
	"github.com/jinzhu/gorm"
	"gopkg.in/vmihailenco/msgpack.v2"
)

//OrmProvider provider of gorm
type OrmProvider struct {
	Db  *gorm.DB
	Enc *gails.Encryptor
}

//Set set
func (p *OrmProvider) Set(k string, v interface{}, f bool) error {
	buf, err := msgpack.Marshal(v)
	if err != nil {
		return err
	}
	if f {
		buf, err = p.Enc.Encode(buf)
		if err != nil {
			return err
		}
	}
	var m Model
	null := p.Db.Where("key = ?", k).First(&m).RecordNotFound()
	m.Key = k
	m.Val = buf
	m.Flag = f
	if null {
		err = p.Db.Create(&m).Error
	} else {
		err = p.Db.Save(&m).Error
	}
	return err
}

//Get get
func (p *OrmProvider) Get(k string, v interface{}) error {
	var m Model
	err := p.Db.Where("key = ?", k).First(&m).Error
	if err != nil {
		return err
	}
	if m.Flag {
		if m.Val, err = p.Enc.Decode(m.Val); err != nil {
			return err
		}
	}
	return msgpack.Unmarshal(m.Val, v)
}
