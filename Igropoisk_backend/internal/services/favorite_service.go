package services

import (
	"igropoisk_backend/internal/repositories"
	"igropoisk_backend/internal/models"
)

type FavoriteService struct {
	favRepo *repositories.FavoriteRepository
}

func NewFavoriteService(favRepo *repositories.FavoriteRepository) *FavoriteService {
	return &FavoriteService{favRepo: favRepo}
}

func (s *FavoriteService) AddFavorite(userID, gameID int) error {
	return s.favRepo.AddFavorite(userID, gameID)
}

func (s *FavoriteService) RemoveFavorite(userID, gameID int) error {
	return s.favRepo.RemoveFavorite(userID, gameID)
}

func (s *FavoriteService) GetFavorites(userID int) ([]models.Game, error) {
	return s.favRepo.GetFavoritesByUser(userID)
}