package routes

import (
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

var db = dbConnection()

func dbConnection() *sqlx.DB {
	return sqlx.MustConnect("mysql", "root:admin@tcp(localhost:3306)/takecare_service")
}
