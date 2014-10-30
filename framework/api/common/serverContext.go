package common

import (
	"fmt"
	. "framework/api/common/types"
	. "framework/api/common/utils"
	db "framework/api/database"
)

//golobal server context variable
var ServerCtx = &ServerContext{}

//server context,will be passed to every handler
//instantiate it in server init function
type ServerContext struct {
	dbSvc db.DataBaseSvcer
}

//program init
func init() {
	initialize()
}
func initialize() {
	var cfg LaserConfig
	//read laser config
	ReadConfig(Lasercfg, &cfg)
	Logging(cfg)
	dbSvr := &db.DataBaseSvc{}
	ServerCtx.dbSvc = dbSvr
	//set the server,user,pwd
	//the database set a placehold,and replace it at a certain time
	ServerCtx.dbSvc.SetConnectionString(fmt.Sprintf("driver={Sql Server};server=%s;database=%s;user id=%s;pwd=%s",
		cfg.Connection.Server,
		"%s", cfg.Connection.User,
		cfg.Connection.Password))
}
