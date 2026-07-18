package repository

import "BookShop/internal/model"

type FavoriteRepository interface {
	GetAll() ([]model.Favorite, error)
	Add(f *model.Favorite) error
	Remove(f *model.Favorite) error
}
