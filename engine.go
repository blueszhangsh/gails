package gails

import "github.com/julienschmidt/httprouter"

//Engine web engine
type Engine interface {
	Mount(*httprouter.Router)
}
