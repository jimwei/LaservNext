package t

import (
	"log"
	_ "net/http"
	_ "os"
	"path/filepath"
	"vendor/toml"
)

const (
	routeConfig = "config/laserRouter.cfg"
)

func ReadConfig() map[string]interface{} {
	p, _ := filepath.Abs(routeConfig)
	log.Println("the path is", p)
	var result = make(map[string]interface{})
	_, err := toml.DecodeFile(p, result)
	if err != nil {
		log.Println("ReadConfig:", err.Error())
		return nil
	} else {
		log.Println(result)
		return result
	}
}
