package repositories

import (
	"database/sql"
	"igropoisk_backend/internal/models"
	
)

type FavoriteRepository struct {
	db *sql.DB
}

func NewFavoriteRepository(db *sql.DB) *FavoriteRepository {
	return &FavoriteRepository{db: db}
}

// Добавить игру в избранное
func (r *FavoriteRepository) AddFavorite(userID, gameID int) error {
	_, err := r.db.Exec(`
		INSERT INTO favorites (user_id, game_id)
		VALUES ($1, $2)
		ON CONFLICT (user_id, game_id) DO NOTHING
	`, userID, gameID)
	return err
}

// Удалить игру из избранного
func (r *FavoriteRepository) RemoveFavorite(userID, gameID int) error {
	_, err := r.db.Exec(`
		DELETE FROM favorites WHERE user_id = $1 AND game_id = $2
	`, userID, gameID)
	return err
}

// Получить все любимые игры пользователя
func (r *FavoriteRepository) GetFavoritesByUser(userID int) ([]models.Game, error) {
	rows, err := r.db.Query(`
		SELECT g.id, g.title, g.description, g.release_date, g.developer, g.publisher, g.cover_url, g.platforms
		FROM favorites f
		JOIN games g ON f.game_id = g.id
		WHERE f.user_id = $1
		ORDER BY f.created_at DESC
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []models.Game
	for rows.Next() {
		var g models.Game
		if err := rows.Scan(
			&g.ID,
			&g.Title,
			&g.Description,
			&g.ReleaseDate,
			&g.Developer,
			&g.Publisher,
			&g.CoverURL,
			&g.Platforms,
		); err != nil {
			return nil, err
		}
		games = append(games, g)
	}
	return games, nil
}
