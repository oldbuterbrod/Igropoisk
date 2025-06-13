package repositories

import (
	"database/sql"
	"errors"
	"igropoisk_backend/internal/models"
	"strings"
)

type ReviewRepository struct {
	db *sql.DB
}

func NewReviewRepository(db *sql.DB) *ReviewRepository {
	return &ReviewRepository{db: db}
}

func (r *ReviewRepository) CreateReview(userID, gameID int, text string) error {
	_, err := r.db.Exec(`
		INSERT INTO reviews (user_id, game_id, text)
		VALUES ($1, $2, $3)
	`, userID, gameID, text)

	if err != nil && strings.Contains(err.Error(), "unique_review_per_user_per_game") {
		return errors.New("вы уже оставили отзыв на эту игру")
	}
	return err
}

func (r *ReviewRepository) GetReviewsByGameID(gameID int) ([]models.Review, error) {
	rows, err := r.db.Query(`
		SELECT r.id, r.user_id, r.game_id, r.text, r.created_at, u.username
		FROM reviews r
		JOIN users u ON r.user_id = u.id
		WHERE r.game_id = $1
		ORDER BY r.created_at DESC
	`, gameID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var r models.Review
		var username string
		if err := rows.Scan(&r.ID, &r.UserID, &r.GameID, &r.Text, &r.CreatedAt, &username); err != nil {
			return nil, err
		}
		r.User = &models.User{ID: r.UserID, Username: username}
		reviews = append(reviews, r)
	}
	return reviews, nil
}

func (r *ReviewRepository) GetReviewsByUserID(userID int) ([]models.Review, error) {
	rows, err := r.db.Query(`
		SELECT r.id, r.user_id, r.game_id, r.text, r.created_at, g.title
		FROM reviews r
		JOIN games g ON r.game_id = g.id
		WHERE r.user_id = $1
		ORDER BY r.created_at DESC
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var r models.Review
		var gameTitle string
		if err := rows.Scan(&r.ID, &r.UserID, &r.GameID, &r.Text, &r.CreatedAt, &gameTitle); err != nil {
			return nil, err
		}
		r.GameTitle = gameTitle 
		reviews = append(reviews, r)
	}
	return reviews, nil
}

func (r *ReviewRepository) DeleteReview(userID, gameID int) error {
	_, err := r.db.Exec(`
		DELETE FROM reviews WHERE user_id = $1 AND game_id = $2
	`, userID, gameID)
	return err
}
