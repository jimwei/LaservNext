package main

import (
	_ "database/sql"
	_ "errors"
	"fmt"
	_ "strconv"
	_ "strings"
	"testing"
)

func Test_ConnectionString(t *testing.T) {
	s := getSqlconnectionString()
	fmt.Println(s)
}
func Test_DB(t *testing.T) {
	db, err := CreateConnection()
	ErrorProcess(err)
	defer db.Close()

	var bumonName string
	err = db.QueryRow("select top 1 Name from Bumon").Scan(&bumonName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("the bumon name is", bumonName)
	var version string
	if err = db.QueryRow("select @@version").Scan(&version); err != nil {
		panic("read version fail.")
	}
	fmt.Println("the sql server version is", version)
}
