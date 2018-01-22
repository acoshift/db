package main

import (
	"database/sql"
	"log"

	"github.com/acoshift/db/postgresql"
)

var settings = "postgres://demouser:demop4ss@demo.upper.io/booktown?sslmode=disable"

type Book struct {
	ID        int    `db:"id"`
	Title     string `db:"title"`
	AuthorID  int    `db:"author_id"`
	SubjectID int    `db:"subject_id"`
}

func main() {
	db, err := sql.Open("postgres", settings)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sess := postgresql.New(db)

	var books []Book
	err = sess.Collection("books").Find().All(&books)
	if err != nil {
		log.Fatal(err)
	}

	for _, book := range books {
		log.Printf("%q (ID: %d)\n", book.Title, book.ID)
	}
}
