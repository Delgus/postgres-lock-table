package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:123456@localhost/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

	lock(db)
	lock(db) //будет висеть и ждать когда закончится первая транзакция
	lock(db)
}

func lock(db *sql.DB){
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	if _, err := tx.Query("SELECT * FROM public.books FOR UPDATE"); err != nil {
		tx.Rollback()
	}
}
