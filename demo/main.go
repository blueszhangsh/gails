package main

import (
	"log"

	"github.com/itpkg/gails"
	_ "github.com/itpkg/gails/engines/cms"
	_ "github.com/itpkg/gails/engines/dict"
	_ "github.com/itpkg/gails/engines/hr"
	_ "github.com/itpkg/gails/engines/note"
	_ "github.com/itpkg/gails/engines/ops"
	_ "github.com/itpkg/gails/engines/team"
)

func main() {
	if err := gails.Main("demo"); err != nil {
		log.Fatal(err)
	}
}
