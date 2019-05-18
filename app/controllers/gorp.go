package controllers

import (
	"github.com/revel/revel"
	"github.com/go-gorp/gorp"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
	"strings"

	"github.com/rl404/tax-calculator/app/models"
)

var Dbm *gorp.DbMap

func InitDB() {
    // db, err := sql.Open("mysql", "HJMG5cEMfF:teogTTIrHe@(remotemysql.com:3306)/HJMG5cEMfF")
	db, err := sql.Open("mysql", getConnectionString())
	if err != nil {
		panic(err)
	}
	Dbm = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	Dbm.AddTableWithName(models.Tax{}, "tax").SetKeys(true, "Id")
	Dbm.AddTableWithName(models.Bill{}, "bill").SetKeys(true, "Id")

	Dbm.CreateTablesIfNotExists()
}

func getParamString(param string, defaultValue string) string {
    p, found := revel.Config.String(param)
    if !found {
        if defaultValue == "" {
            log.Fatal("Cound not find parameter: " + param)
        } else {
            return defaultValue
        }
    }
    return p
}

func getConnectionString() string {
    host := getParamString("db.host", "")
    port := getParamString("db.port", "3306")
    user := getParamString("db.user", "")
    pass := getParamString("db.password", "")
    dbname := getParamString("db.name", "")
    protocol := getParamString("db.protocol", "tcp")
    dbargs := getParamString("dbargs", " ")
    timezone := getParamString("db.timezone", "parseTime=true&loc=Asia%2FJakarta")

    if strings.Trim(dbargs, " ") != "" {
        dbargs = "?" + dbargs
    } else {
        dbargs = ""
    }
    return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s?%s", user, pass, protocol, host, port, dbname, dbargs, timezone)
}