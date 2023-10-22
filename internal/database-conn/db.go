package databaseconn

import (
	"database/sql"
	"log"

	"github.com/RazvanBerbece/UnifyFootballBot/internal/globals"
	"github.com/go-sql-driver/mysql"
)

type Database struct {
	cfg mysql.Config
	Db  *sql.DB
}

func (d *Database) ConnectDatabaseHandle() {

	db, err := sql.Open("mysql", globals.MySqlConnectionString)
	if err != nil {
		log.Fatal("Connection to database cannot be established :", err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal("Database cannot be reached :", pingErr)
	}

	d.Db = db

}
