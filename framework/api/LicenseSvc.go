package laserframework

import (
	"net/http"
	"time"
)

type Service struct {
	ProductID    int       `json:"productid"`
	ProductName  string    `json:"productname"`
	ServiceID    int       `json:"serviceid"`
	ServiceName  string    `json:"servicename"`
	ServiceType  string    `json:"servicetype"`
	ServiceCount int       `json:"servicecount"`
	StartDate    time.Time `json:"startdate"`
	EndDate      time.Time `json:"enddate"`
}
type Account struct {
	AccountID   int       `json:"accountid"`
	AccountName string    `json:"accountname"`
	Services    []Service `json:"services"`
}
type LicenseFile struct {
	AccountId int    `json:"accountid"`
	License   string `json:"license"`
}
type AccountSvc struct{}

//create license file
func (self AccountSvc) CreateLicenseFile() (string, error) {
	return "", nil
}

//load license file from db
func (self AccountSvc) GetLicenseFile(accountID int) (string, error) {
	return "", nil
}

//valid the license file whether valid or not
func (self AccountSvc) ValidLicenseFile(license string) bool {
	return true
}
