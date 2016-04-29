package main

import (
	"log"

	"github.com/itpkg/gails"
	_ "github.com/itpkg/gails/engines/base"
)

func main() {
	if err := gails.Main(); err != nil {
		log.Fatal(err)
	}
}
