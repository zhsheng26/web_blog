package support

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

func ConnectSQL(host, port, usr, pwd, dbName string) (db *DB, err error) {
	dbSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		usr,
		pwd,
		host,
		port,
		dbName,
	)

	mysqlDb, err := sql.Open("mysql", dbSource)
	if err != nil {
		panic(err)
	}
	dbConn.SQL = mysqlDb
	return dbConn, err
}
