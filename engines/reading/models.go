package reading

import (
	"github.com/itpkg/gails"
	"github.com/itpkg/gails/engines/auth"
)

type Note struct {
	gails.Model
	Title  string `gorm:"not null;index;type:VARCHAR(255)" json:"title"`
	Body   string `gorm:"not null;type:TEXT" json:"body"`
	UserID uint   `gorm:"not null" json:"userId"`
	User   auth.User
}

func (Note) TableName() string {
	return "reading_notes"
}
