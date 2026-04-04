package repository

import "BookShop/internal/model"

type BookRepository interface {
	GetAll() ([]model.Book, error)
	GetByID(id int) (*model.Book, error)
	Create(book *model.Book) error
	Update(book *model.Book) error
	Delete(id int) error
}


