package db

import (
	"database/sql"
	"log"

	"khakibee/site-api/config"

	//mysql loading driver
	_ "github.com/go-sql-driver/mysql"
)

// New :creates new db connect
func New() *sql.DB {
	// Open database connection
	db, err := sql.Open("mysql", config.Conf.Conn)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxIdleConns(0)   // this is the root problem! set it to 0 to remove all idle connections
	db.SetMaxOpenConns(100) // or whatever is appropriate for your setup.

	// Check if connection is ok
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connected")

	return db
}