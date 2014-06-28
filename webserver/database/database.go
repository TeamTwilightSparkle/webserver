package database

import (
	"fmt"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/coopernurse/gorp"
	"os"
)

var DBMap *gorp.DbMap
var db *sql.DB

func BootDatabase(dbuser, dbname string) {
	var err error
	db, err = sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s sslmode=disable", dbuser, dbname))
	if err != nil {
		fmt.Printf("Failied to initialize Database: %s\n", dbname)
		os.Exit(2)
	}
	DBMap = &gorp.DbMap{Db: db}
	fmt.Printf("Initialized PostgreSQL User: %s @ Database: %s\n", dbuser, dbname)
}

func CloseDB() {
	db.Close()
}
