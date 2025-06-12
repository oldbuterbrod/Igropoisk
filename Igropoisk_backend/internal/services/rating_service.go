package services

import "igropoisk_backend/internal/repositories"

type RatingService struct {
	repo *repositories.RatingRepository
}

func NewRatingService(repo *repositories.RatingRepository) *RatingService {
	return &RatingService{repo: repo}
}

func (s *RatingService) RateGame(userID, gameID, score int) error {
	return s.repo.SetRating(userID, gameID, score)
}

func (s *RatingService) GetUserRating(userID, gameID int) (int, error) {
	return s.repo.GetRating(userID, gameID)
}

func (s *RatingService) GetAverageRating(gameID int) (float64, error) {
	return s.repo.GetAverageRating(gameID)
}

func (s *RatingService) GetRatingsByUser(userID int) (map[int]repositories.RatingInfo, error) {
	return s.repo.GetRatingsByUser(userID)
}

func (s *RatingService) DeleteRating(userID, gameID int) error {
	return s.repo.DeleteRating(userID, gameID)
}
