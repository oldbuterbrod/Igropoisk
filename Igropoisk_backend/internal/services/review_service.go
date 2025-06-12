package services

import (
	"igropoisk_backend/internal/models"
	"igropoisk_backend/internal/repositories"
)

type ReviewService struct {
	repo *repositories.ReviewRepository
}

func NewReviewService(repo *repositories.ReviewRepository) *ReviewService {
	return &ReviewService{repo: repo}
}

func (s *ReviewService) CreateReview(userID, gameID int, text string) error {
	return s.repo.CreateReview(userID, gameID, text)
}

func (s *ReviewService) GetReviewsByGameID(gameID int) ([]models.Review, error) {
	return s.repo.GetReviewsByGameID(gameID)
}

func (s *ReviewService) GetReviewsByUserID(userID int) ([]models.Review, error) {
	return s.repo.GetReviewsByUserID(userID)
}

func (s *ReviewService) DeleteReview(userID, gameID int) error {
	return s.repo.DeleteReview(userID, gameID)
}
