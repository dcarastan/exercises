package main

import (
	"fmt"
	"log"
	"net/http"

	// "database/sql"
	sql "github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE IF NOT EXISTS books (
	title text,
	author text
)`

// ShowBooks func
func ShowBooks(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var title, author string
		err := db.QueryRow("select title, author from books").Scan(&title, &author)
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(rw, "The first book is '%s' by '%s'", title, author)
	})
}

// NewDB func
func NewDB() *sql.DB {
	log.Println("Opening DB")
	db, err := sql.Open("sqlite3", "database.sqlite")
	if err != nil {
		panic(err)
	}

	log.Println("Creating table")
	_, err = db.Exec(schema)
	if err != nil {
		panic(err)
	}

	log.Println("Inserting books")
	_, err = db.Exec("INSERT INTO books (title, author) VALUES ('Amintiri din copilarie', 'Ion Creanga')")
	if err != nil {
		panic(err)
	}

	return db
}

func main() {
	log.Println("Start")
	db := NewDB()
	log.Println("Listening on :8888")
	http.ListenAndServe(":8888", ShowBooks(db))
}
