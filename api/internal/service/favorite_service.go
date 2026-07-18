package service

import (
	"BookShop/internal/model"
	"BookShop/internal/repository"
)

type FavoriteService struct {
	repo repository.FavoriteRepository
}

func InitFavoriteService(repo repository.FavoriteRepository) *FavoriteService {
	return &FavoriteService{repo: repo}
}

func (s *FavoriteService) GetAll() ([]model.Favorite, error) {
	return s.repo.GetAll()
}

func (s *FavoriteService) Add(f *model.Favorite) (string, error) { return s.repo.Add(f) }
