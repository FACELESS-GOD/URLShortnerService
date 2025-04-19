package Utility

import (
	"database/sql"

	Config "github.com/FACELESS-GOD/URLShortnerService.git/Helper/MetaData"
	_ "github.com/go-sql-driver/mysql"
)

var DbInstance *sql.DB

func InitiateDatabase() {
	db, err := sql.Open("mysql", Config.DbConnString)
	if err != nil {
		panic(err)
	}
	DbInstance = db

}
