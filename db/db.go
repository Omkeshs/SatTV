package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	MYSQLDBUSER     = "user"
	MYSQLDBPASS     = "password"
	MYSQLDBNAME     = "d2h"
	TESTMYSQLDBNAME = "d2h-test"
)

//NewDbConnection for client connection
func NewDbConnection() *gorm.DB {

	strConnection := MYSQLDBUSER + ":" + MYSQLDBPASS + "@/" + MYSQLDBNAME + "?charset=utf8&parseTime=True&loc=Local"

	//Open Connection for GORM instance
	client, err := gorm.Open("mysql", strConnection)

	if err != nil {
		panic("Error! Can't create connection to database")
	}
	return client

}

//NewTestDbConnection for client connection
func NewTestDbConnection() *gorm.DB {

	strConnection := MYSQLDBUSER + ":" + MYSQLDBPASS + "@/" + TESTMYSQLDBNAME + "?charset=utf8&parseTime=True&loc=Local"

	//Open Connection for GORM instance
	client, err := gorm.Open("mysql", strConnection)

	if err != nil {
		panic("Error! Can't create connection to database")
	}
	return client

}
