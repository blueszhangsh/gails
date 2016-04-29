package main

import (
	"log"

	"github.com/itpkg/gails"
	_ "github.com/itpkg/gails/engines/base"
)

func main() {
	if err := gails.Main("demo"); err != nil {
		log.Fatal(err)
	}
}
