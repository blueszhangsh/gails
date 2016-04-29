package gails_test

import (
	"testing"

	"github.com/itpkg/gails"
)

type M1 struct {
	S string `inject:"hello"`
	I int    `inject:"version"`
}

type M2 struct {
	M *M1 `inject:""`
}

func TestInject(t *testing.T) {
	gails.Use(&M2{})
	gails.Map("hello", "Hello, Gails!")
	gails.Map("version", 2016)
	if err := gails.Run(func(m1 *M1, m2 *M2) {
		t.Logf("M1: %+v", m1)
		t.Logf("M2: %+v", m2)
	}); err != nil {
		t.Fatal(err)
	}
}
