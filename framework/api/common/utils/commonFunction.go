package common

import (
	"log"
	"path/filepath"
	"reflect"
	"toml"
)

//check the v passed whether is nil or not
func IsNil(v interface{}) bool {
	return (v == nil)
}

//check the v passed whether is not nil or not
func IsNotNil(v interface{}) bool {
	return !IsNil(v)
}

//now only write to console
//will be ouptput to a log file later
func Logging(messsage interface{}) {
	log.Println(messsage)
}

//read config function
func ReadConfig(cfgName string, desStruct interface{}) interface{} {
	//read laser config
	cfgPath, _ := filepath.Abs(cfgName)
	//check whether the deststruct is addr or not
	var isAddr = (reflect.ValueOf(desStruct).Kind() == reflect.Ptr)

	if IsNotNil(desStruct) {
		if isAddr {
			toml.DecodeFile(cfgPath, desStruct)
			return desStruct
		} else {
			toml.DecodeFile(cfgPath, &desStruct)
			return desStruct
		}

	} else {
		var dest map[string]interface{}
		toml.DecodeFile(cfgPath, &dest)
		return &dest
	}

}
