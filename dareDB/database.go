package main

import (
	"dareAPI/api"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = ""
	password = ""
	dbname   = ""
)

func createDatabase() {
	// The only difference between Sprintf() and Printf() is that Sprintf writes data to a character array,
	// while Printf() writes data to stdout, the standard output device
	connInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)
	// Open opens a dareDB specified by its dareDB driver name and a driver-specific data source name
	// func Open(driverName, dataSourceName string) (*DB, error)
	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		log.Fatal(err)
	}

	// Exec executes a query without returning any rows.
	// func (db *DB) Exec(query string, args ...any) (Result, error)
	// Create dareDB
	_, err = db.Exec("CREATE DATABASE " + dbname)

	if err != nil {
		//handle the error
		log.Fatal(err)
	}

	// Create Table
	_, err = db.Exec("CREATE TABLE daresdb (DareId integer,DareQuestion varchar(500) )")
	if err != nil {
		log.Fatal(err)
	}

}

func InsertDare(d *api.Dare) (dareID int) {
	connInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// open dareDB
	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	insertQuery := "INSERT INTO daresdb(DareId = ?, DareQuestion = ?)"
	_, err = db.Exec(insertQuery, d.DareID, d.DareQuestion)
	if err != nil {
		log.Fatal(err)
	}

	return d.DareID
}

func main() {
	createDatabase()
	testDare := api.Dare{
		DareID:       1,
		DareQuestion: "This is a test question.",
	}

	InsertDare(&testDare)
}
