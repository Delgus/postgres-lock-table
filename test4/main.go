package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:123456@localhost/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

	lock(db)
	lock(db)
	lock(db)


}

func lock(db *sql.DB){
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	var lock bool
	if err := tx.QueryRow("select pg_try_advisory_xact_lock(1)").Scan(&lock); err != nil {
		tx.Rollback()
	}
	if !lock {
		fmt.Println("i can't lock table")
	}
}
