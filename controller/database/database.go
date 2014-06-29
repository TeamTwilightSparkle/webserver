package database

import (
	"fmt"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/coopernurse/gorp"
	"os"
)

type connection struct {
	dbmap		*gorp.DbMap
	db			*sql.DB
}

var Conn *connection = new(connection)

func BootDatabase(dbuser, dbname string) {
	var err error

	Conn.db, err = sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s sslmode=disable", dbuser, dbname))
	if err != nil {
		fmt.Printf("Failied to initialize Database: %s\n", dbname)
		os.Exit(2)
	}
	Conn.dbmap = &gorp.DbMap{Db: Conn.db}
	fmt.Printf("Initialized PostgreSQL User: %s @ Database: %s\n", dbuser, dbname)
}

func CloseDB() {
	fmt.Println("Closing Database")
	Conn.db.Close()
}

func (conn *connection) Select(i interface {}, q string, args ...interface {}) error {
	if _, err := Conn.dbmap.Select(i, q, args...); err != nil {
		return err
	}
	return nil
}

func (conn *connection) SelectInt(q string) (i int64, err error) {
	return Conn.dbmap.SelectInt(q)
}
