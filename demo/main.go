package main

import (
	"log"

	"github.com/itpkg/gails"
	_ "github.com/itpkg/gails/engines/cms"
	_ "github.com/itpkg/gails/engines/hr"
	_ "github.com/itpkg/gails/engines/ops"
	_ "github.com/itpkg/gails/engines/reading"
	_ "github.com/itpkg/gails/engines/site"
	_ "github.com/itpkg/gails/engines/team"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	if err := gails.Main("demo"); err != nil {
		log.Fatal(err)
	}
}
