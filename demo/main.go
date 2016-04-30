package main

import (
	"log"

	"github.com/itpkg/gails"
	_ "github.com/itpkg/gails/engines/auth"
	_ "github.com/itpkg/gails/engines/base"
	_ "github.com/itpkg/gails/engines/cms"
	_ "github.com/itpkg/gails/engines/hr"
	_ "github.com/itpkg/gails/engines/ops"
	_ "github.com/itpkg/gails/engines/reading"
	_ "github.com/itpkg/gails/engines/team"
)

func main() {
	if err := gails.Main("demo"); err != nil {
		log.Fatal(err)
	}
}
