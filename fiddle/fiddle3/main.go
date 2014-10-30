//test toml format,and can decode in a map[string]interface{}
package main

import (
	"log"
	"net/http"
	"vendor/toml"
)

const (
	routeConfig = "config/laserRouter.cfg"
)

func main() {
	log.Println(http.Dir(routeConfig))
	var result = make(map[string]interface{})
	_, err := toml.DecodeFile(routeConfig, result)
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("the config is", result)
		var logonHanlderName = result["route"].(map[string]interface{})["logon"].(string)
		log.Println("the logon handler is", logonHanlderName)
	}
}
