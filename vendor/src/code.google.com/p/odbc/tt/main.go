package main

import (
	"database/sql"
	//"errors"
	"flag"
	"fmt"
	//"runtime"
	//"strconv"
	//"strings"
	//"sync/atomic"
	//"testing"
	//"time"
	"bytes"
	_ "code.google.com/p/odbc"
	_ "code.google.com/p/odbc/api"
)

var (
	mssrv    = flag.String("mssrv", "xa-lsr-jimweiw7", "server name")
	msdb     = flag.String("msdb", "Account_001_Kaikei", "database")
	msuser   = flag.String("user", "sa", "user name")
	mspwd    = flag.String("pwd", "xA123456", "password")
	msdriver = flag.String("msdriver", "sql server", "msdriver")
	msport   = flag.String("msport", "1433", "connect port")
)

//创建数据库连接字符串
func CreateConnection() (db *sql.DB, err error) {
	sqlConnString := getSqlconnectionString()
	db, err = sql.Open("odbc", sqlConnString)
	ErrorProcess(err)
	return db, nil
}

//构造连接字符串
func getSqlconnectionString() string {
	buf := bytes.NewBufferString("")
	buf.WriteString("Driver=%s;")
	buf.WriteString("server=%s;")
	buf.WriteString("database=%s;")
	buf.WriteString("user id=%s;")
	buf.WriteString("pwd=%s;")

	connString := fmt.Sprintf(buf.String(), *msdriver, *mssrv, *msdb, *msuser, *mspwd)
	return connString
}
func main() {

}
