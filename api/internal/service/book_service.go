package service

import (
	"BookShop/internal/model"
	"BookShop/internal/repository"
)

type BookService struct {
	repo repository.BookRepository
}

func InitBookService(repo repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) GetAll() ([]model.Book, error) {
	return s.repo.GetAll()
}

func (s *BookService) GetByID(id int) (*model.Book, error) {
	return s.repo.GetByID(id)
}

func (s *BookService) Create(book *model.Book) error {
	return s.repo.Create(book)
}

func (s *BookService) Update(book *model.Book) error {
	return s.repo.Update(book)
}

func (s *BookService) Delete(id int) error {
	return s.repo.Delete(id)
}
