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

//Loop loop beans
func Loop(f func(string, interface{}) error) error {
	for _, o := range beans.Objects() {
		if e := f(o.Name, o.Value); e != nil {
			return e
		}
	}
	return nil
}

//Init build beans
func Init() error {
	return beans.Populate()
}

//Run main entry
func Run(f interface{}) (interface{}, error) {
	if reflect.TypeOf(f).Kind() != reflect.Func {
		return nil, errors.New("bad type, need a func")
	}

	return nil, nil
}
