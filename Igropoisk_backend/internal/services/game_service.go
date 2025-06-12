package services

import (
	"igropoisk_backend/internal/models"
	"igropoisk_backend/internal/repositories"
)

type GameService struct {
	repo *repositories.GameRepository
}

func NewGameService(repo *repositories.GameRepository) *GameService {
	return &GameService{repo: repo}
}

func (s *GameService) GetGames() ([]models.Game, error) {
	return s.repo.GetAll()
}

func (s *GameService) GetGameByID(id int) (*models.Game, error) {
	return s.repo.GetGameByID(id)
}

func (s *GameService) SearchGames(query string) ([]models.Game, error) {
	return s.repo.SearchGames(query)
}

func (s *GameService) GetGenresByGameID(gameID int) ([]string, error) {
	return s.repo.GetGenresByGameID(gameID)
}

func (s *GameService) GetSimilarGames(gameID int) ([]models.Game, error) {
	return s.repo.GetSimilarGames(gameID)
}

func (s *GameService) GetRecommendedGames(userID int) ([]models.Game, error) {
	return s.repo.GetRecommendedGamesByUser(userID)
}

func (s *GameService) GetTopRatedGames() ([]models.Game, error) {
	return s.repo.GetTopRatedGames(5)
}

func (s *GameService) GetUpcomingGames() ([]models.Game, error) {
	return s.repo.GetUpcomingGames(3)
}

func (s *GameService) GetRecentGames() ([]models.Game, error) {
	return s.repo.GetRecentGames(3)
}

func (s *GameService) GetPopularGames() ([]models.Game, error) {
	return s.repo.GetPopularGames(3)
}
