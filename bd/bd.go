package bd

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func BdconMysql() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "shoplek"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	if err != nil {
		return nil, err
	}

	// Ensure the database connection is valid
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
