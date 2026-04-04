package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://default:secret@localhost:5432/default?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	books := []struct {
		Title, Author string
		Price         float64
	}{
		{"The Go Programming Language", "Alan Donovan", 29.99},
		{"Clean Code", "Robert Martin", 24.99},
		{"Domain-Driven Design", "Eric Evans", 34.99},
	}

	for _, b := range books {
		_, err := db.Exec(`INSERT INTO books (title, author, price) VALUES ($1, $2, $3)`, b.Title, b.Author, b.Price)
		if err != nil {
			log.Printf("failed to insert %s: %v", b.Title, err)
		}
	}
	log.Println("Seeding complete")
}
