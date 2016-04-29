package gails

import (
	"errors"
	"reflect"

	"github.com/facebookgo/inject"
)

var beans inject.Graph

//Use inject objects without name
func Use(args ...interface{}) {
	for _, o := range args {
		beans.Provide(&inject.Object{Value: o})
	}
}

//Map inject object with name
func Map(n string, o interface{}) {
	beans.Provide(&inject.Object{Value: o, Name: n})
}

//Run main entry
func Run(f interface{}) error {
	if reflect.TypeOf(f).Kind() != reflect.Func {
		return errors.New("bad type, need a func")
	}

	if err := beans.Populate(); err != nil {
		return err
	}
	return nil
}
