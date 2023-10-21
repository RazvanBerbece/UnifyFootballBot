package databaseconn

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

type Database struct {
	cfg mysql.Config
}

func (d *Database) ConfigureConn() {
	d.cfg = mysql.Config{
		User:   os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_PASSWORD"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: os.Getenv("MYSQL_DATABASE"),
	}
}

func (d *Database) GetDatabaseHandle() *sql.DB {

	db, err := sql.Open("mysql", d.cfg.FormatDSN())
	if err != nil {
		log.Fatal("Connection to database cannot be established :", err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal("Database cannot be reached :", pingErr)
	}

	return db
}
