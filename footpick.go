package main

import (
	"database/sql"
	"log"

	"github.com/daveseo901/footpick/league"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

func main() {
	// TODO: hook up database (do we even need one?)
	dbFile := "league.db"
	db, err := sql.Open("sqlite3", dbFile)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Populating db")
	err = league.PopulateDatabase(db, "/home/dseo/Downloads/maddennfl24ratings.csv")
	if err != nil {
		log.Fatal(err)
	}
	return
}
