package gails

import (
	"errors"
	"fmt"
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
func Run(f interface{}) ([]interface{}, error) {
	ft := reflect.TypeOf(f)
	if ft.Kind() != reflect.Func {
		return nil, errors.New("bad type, need a func")
	}
	fv := reflect.ValueOf(f)
	args := make([]reflect.Value, 0)
	for i := 0; i < ft.NumIn(); i++ {
		fet := ft.In(i)
		for _, bn := range beans.Objects() {
			if fet == reflect.TypeOf(bn.Value) {
				args = append(args, reflect.ValueOf(bn.Value))
				break
			}
		}
		if len(args) != i+1 {
			return nil, fmt.Errorf("can't find type %s.%s", fet.PkgPath(), fet.Name())
		}
	}

	res := fv.Call(args)
	ret := make([]interface{}, 0)
	for _, v := range res {
		ret = append(ret, v.Interface())
	}
	return ret, nil
}
