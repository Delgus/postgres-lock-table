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

	if _, err := tx.Exec("LOCK TABLE public.books NOWAIT"); err != nil {
		tx.Rollback()
		fmt.Println(`i can't lock table!!!`)
	}
}
