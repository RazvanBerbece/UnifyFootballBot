package databaseconn

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

type Database struct {
	cfg mysql.Config
	Db  *sql.DB
}

func (d *Database) ConnectDatabaseHandle() {

	db, err := sql.Open("mysql", os.Getenv("UNIFYFOOTBALL_DB_CONNSTRING"))
	if err != nil {
		log.Fatal("Connection to database cannot be established :", err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal("Database cannot be reached :", pingErr)
	}

	d.Db = db

}
