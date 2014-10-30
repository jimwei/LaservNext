package main

import (
	"fmt"
	s "leyservNext/fiddle/fiddle2"
	"reflect"
)

func init() {
}
func main() {
	fmt.Println("hello guy.")
	var per s.Person
	result := reflect.ValueOf(&per).MethodByName("GetName").Call([]reflect.Value{})
	fmt.Println(result)
	reflect.ty
	//structName := "fiddle2.Person"
	//funcName := "GetName"
	//sv := reflect.ValueOf(structName)
	//md := sv.MethodByName(funcName)
	//v := md.Call([]reflect.Value{})
	//fmt.Println(v)

}
