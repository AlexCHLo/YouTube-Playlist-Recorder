package myfunctions
// SQLite Function Implementation


import (
	"fmt"
	// "database/sql"
	// _ "github.com/mattn/go-sqlite3"
)


// Create a database. init bruv
func databaseInit() {
	fmt.Println("Sanity Check - databaseInit")

	file, err := os.Create("database.db")

	if err != nil {
		log.Fatal(err)
	}

	file.Close()

}

// 
