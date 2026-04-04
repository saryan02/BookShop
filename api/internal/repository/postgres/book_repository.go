package postgres

import (
	"database/sql"

	"BookShop/internal/model"
)

type BookRepository struct {
	db *sql.DB
}

func InitBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) GetAll() ([]model.Book, error) {
	rows, err := r.db.Query(`SELECT id, title, author, price FROM books`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []model.Book
	for rows.Next() {
		var b model.Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Price); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, rows.Err()
}

func (r *BookRepository) GetByID(id int) (*model.Book, error) {
	var b model.Book
	err := r.db.QueryRow(`SELECT id, title, author, price FROM books WHERE id = $1`, id).
		Scan(&b.ID, &b.Title, &b.Author, &b.Price)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func (r *BookRepository) Create(book *model.Book) error {
	return r.db.QueryRow(
		`INSERT INTO books (title, author, price) VALUES ($1, $2, $3) RETURNING id`,
		book.Title, book.Author, book.Price,
	).Scan(&book.ID)
}

func (r *BookRepository) Update(book *model.Book) error {
	_, err := r.db.Exec(
		`UPDATE books SET title=$1, author=$2, price=$3 WHERE id=$4`,
		book.Title, book.Author, book.Price, book.ID,
	)
	return err
}

func (r *BookRepository) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM books WHERE id = $1`, id)
	return err
}
