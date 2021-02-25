package config

import (
	"database/sql"

	"github.com/go-sql-driver-mysql"
)

func GetDB() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "db_user"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return

}
