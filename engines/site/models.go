package site

import "github.com/itpkg/gails"

//Notice notice model
type Notice struct {
	gails.Model
	Lang    string `gorm:"not null;type:varchar(8);index" json:"lang"`
	Content string `gorm:"not null;type:text" json:"content"`
}
