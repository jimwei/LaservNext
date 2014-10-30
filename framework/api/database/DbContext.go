// database package
package database

import (
	_ "code.google.com/p/odbc"
	. "framework/api/common/utils"
	"github.com/jmoiron/sqlx"
)

// database service interface
type DataBaseSvcer interface {
	ConnectionString() string
	SetConnectionString(connstring string)
	CreateDb() *sqlx.DB
}

//database service
type DataBaseSvc struct {
	connectionString string
}

//set the connection string
func (this *DataBaseSvc) SetConnectionString(connstring string) {
	this.connectionString = connstring

}

//create db provider
func (this *DataBaseSvc) CreateDb() *sqlx.DB {
	db, err := sqlx.Open("odbc", this.connectionString)
	if IsNotNil(err) {
		Logging(err.Error())
	}

	return db
}

//get the connection string
func (this *DataBaseSvc) ConnectionString() string {
	return this.connectionString
}
