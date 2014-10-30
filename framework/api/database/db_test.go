package database

import (
	"log"
	"testing"
)

func Test_db(t *testing.T) {
	dbSvc := new(DataBaseSvc)
	dbSvc.SetConnectionString("driver={SQL Server};server=xa-lsr-jimweiw7\\sqlserver14;database=ACCOUNT_000001_KAIKEI;user id=sa;pwd=xA123456;")
	db := dbSvc.CreateDb()
	var count int
	row := db.QueryRow("select count(*) from bumon")
	err := row.Scan(&count)
	if err != nil {
		log.Println("fetch the counts fail." + err.Error())
	}
	log.Println("the bumon count is", count)
}
func Test_SqlxDB(t *testing.T) {
	dbSvc := &DataBaseSvc{}
	dbSvc.SetConnectionString("driver={SQL Server};server=xa-lsr-jimweiw7\\sqlserver14;database=ACCOUNT_000001_KAIKEI;user id=sa;pwd=xA123456;")
	db := dbSvc.CreateDb()
	var count int
	row := db.QueryRow("select count(*) from bumon")
	err := row.Scan(&count)
	if err != nil {
		log.Println("fetch the counts fail." + err.Error())
	}
	log.Println("the bumon count is", count)

}
