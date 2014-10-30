package Framework

import (
	"framework/api/common"
	"net/http"
)

//router interface
type IRouter interface {
	Handler() http.Handler
	SetHanler(http.Handler)
	Execute() (bool error)
}

//router struct
type AbstractRouter struct {
	handler http.Handler
}

//get the http handler
func (this *AbstractRouter) Hanlder() http.Handler {
	return this.handler
}

//set the http hanlder
func (this *AbstractRouter) SetHandler(h http.Handler) {
	this.handler = h
}

//this function should must be implemented in every product router struct.
//this router is a abstract struct
func (this *AbstractRouter) Execute() (bool, error) {
	return false, NotImplementedException
}
