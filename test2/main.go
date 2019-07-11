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

func lock(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	var isbn, title string
	if err := tx.QueryRow(`SELECT isbn,title FROM public.books WHERE title != 'Petya' LIMIT 1 FOR UPDATE`).Scan(&isbn, &title); err != nil {
		fmt.Printf("query %v\n", err)
		tx.Rollback()
		return
	}
	fmt.Printf("name before: %s\n", title)
	if _, err := tx.Exec(`UPDATE public.books SET title='Petya' where isbn=$1`, isbn); err != nil {
		fmt.Println(err)
		tx.Rollback()
	}
	tx.Commit()
}
