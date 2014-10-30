package Framework

import (
	_ "net/http"
)

//router dispatch interface
//in the
type IRouterDispatcher interface {
	Dispatch(productName string, routeName string) IRouter
}

//route dispatcher interface,inherit iouterdispatcher interface
//when a reqest arrived at the server side,the main router dispatcher will dispatch the request to a product router dispatcher
// which configed, and the product router dispatcher will return a http.Handler to main router dispatcher to invoke.
type ILeyserRouterDispatcher interface {
	IRouterDispatcher
	//in the chain,may include kaiker router dispatcher,shisan router dispatcher,gakuhi router dispatcher,etc.
	// get router dispatcher list function
	RouterDispatcherList() []IRouter
	//add a router dispatcher into the inner router dispatcher list
	//set router dispatcher function
	//return :the router dispatcher list
	AddRouterDispatcher(IRouter) []IRouter
}
type RouterDispatcher struct{}
