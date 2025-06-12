package repositories

import (
	"database/sql"
)

type RatingRepository struct {
	db *sql.DB
}

func NewRatingRepository(db *sql.DB) *RatingRepository {
	return &RatingRepository{db: db}
}

func (r *RatingRepository) SetRating(userID, gameID, score int) error {
	query := `
		INSERT INTO ratings (user_id, game_id, score)
		VALUES ($1, $2, $3)
		ON CONFLICT (user_id, game_id)
		DO UPDATE SET score = EXCLUDED.score;
	`
	_, err := r.db.Exec(query, userID, gameID, score)
	return err
}

func (r *RatingRepository) GetRating(userID, gameID int) (int, error) {
	var score int
	err := r.db.QueryRow(`SELECT score FROM ratings WHERE user_id=$1 AND game_id=$2`, userID, gameID).Scan(&score)
	return score, err
}

func (r *RatingRepository) GetAverageRating(gameID int) (float64, error) {
	var avg float64
	err := r.db.QueryRow(`SELECT AVG(score) FROM ratings WHERE game_id=$1`, gameID).Scan(&avg)
	return avg, err
}

// Новая структура для фронта: score + title
type RatingInfo struct {
	Score int    `json:"score"`
	Title string `json:"title"`
}

func (r *RatingRepository) GetRatingsByUser(userID int) (map[int]RatingInfo, error) {
	rows, err := r.db.Query(`
		SELECT r.game_id, r.score, g.title
		FROM ratings r
		JOIN games g ON r.game_id = g.id
		WHERE r.user_id = $1
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ratings := make(map[int]RatingInfo)
	for rows.Next() {
		var gameID, score int
		var title string
		if err := rows.Scan(&gameID, &score, &title); err != nil {
			return nil, err
		}
		ratings[gameID] = RatingInfo{
			Score: score,
			Title: title,
		}
	}
	return ratings, nil
}

func (r *RatingRepository) DeleteRating(userID, gameID int) error {
	_, err := r.db.Exec(`DELETE FROM ratings WHERE user_id = $1 AND game_id = $2`, userID, gameID)
	return err
}
