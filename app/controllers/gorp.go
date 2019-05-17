package controllers

import (
	"github.com/go-gorp/gorp"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/rl404/tax-calculator/app/models"
)

var Dbm *gorp.DbMap

func InitDB() {
	db, err := sql.Open("mysql", "HJMG5cEMfF:teogTTIrHe@(remotemysql.com:3306)/HJMG5cEMfF")
	if err != nil {
		panic(err)
	}
	Dbm = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	Dbm.AddTableWithName(models.Tax{}, "tax").SetKeys(true, "Id")
	Dbm.AddTableWithName(models.Bill{}, "bill").SetKeys(true, "Id")

	Dbm.CreateTablesIfNotExists()
}