package gails_test

import (
	"testing"

	"github.com/itpkg/gails"
)

type M1 struct {
	S1 string `inject:"hello"`
	I1 int    `inject:"version"`
}

type M2 struct {
	M *M1 `inject:""`
}

type M3 struct {
	M1
	S3 string `inject:"hello"`
	I3 int    `inject:"version"`
}

func TestInject(t *testing.T) {
	gails.Use(&M2{}, &M3{})
	gails.Map("hello", "Hello, Gails!")
	gails.Map("version", 2016)

	if err := gails.Init(); err != nil {
		t.Fatal(err)
	}

	gails.Loop(func(n string, o interface{}) error {
		t.Logf("%s: %+v", n, o)
		return nil
	})
	if val, err := gails.Run(func(m1 *M1, m2 *M2, m3 *M3) {
		t.Logf("M1: %+v", m1)
		t.Logf("M2: %+v", m2)
		t.Logf("M3: %+v", m3)
	}); err == nil {
		t.Logf("%+v", val)
	} else {
		t.Fatal(err)
	}
}
