package i18n_test

import (
	"testing"

	"github.com/itpkg/gails/i18n"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/op/go-logging"
	"golang.org/x/text/language"
)

var lang = &language.SimplifiedChinese

//var logger = log.New(os.Stdout, "[test]", 0)
var logger = logging.MustGetLogger("example")

func TestOrm(t *testing.T) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		t.Fatal(err)
	}
	db.LogMode(true)
	db.AutoMigrate(&i18n.Locale{})
	testProvider(t, &i18n.OrmProvider{Db: db, Logger: logger})
}

func testProvider(t *testing.T, p i18n.Provider) {
	key := "hello"
	val := "你好"
	p.Set(lang, key, val)
	p.Set(lang, key+".1", val)
	if val1 := p.Get(lang, key); val != val1 {
		t.Errorf("want %s, get %s", val, val1)
	}
	ks := p.Keys(lang)
	if len(ks) == 0 {
		t.Errorf("empty keys")
	} else {
		t.Log(ks)
	}
	p.Del(lang, key)
}
