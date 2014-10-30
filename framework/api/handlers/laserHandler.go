package Framework

import (
	. "framework/api/common"
	"net/http"
)

//laser framework handler,inherit http.handler
type ILaserHandler interface {
	http.Handler
	ServerContext() *ServerContext
	SetServerContext(*ServerContext)
}

//framework handler
type LaserHandler struct {
	serverCtx *ServerContext
}

//implement the http.handler,
func (this *LaserHandler) ServeHttp(w http.ResponseWriter, r *http.Request) {
	//todo: framework handler route dispatcher
}

//get server context object
func (this *LaserHandler) ServerContext() *ServerContext {
	return this.serverCtx
}

//set the server context object
func (this *LaserHandler) SetServerContext(serverCtx *ServerContext) {
	this.serverCtx = serverCtx
}
