package repositories

import (
	"database/sql"
	"igropoisk_backend/internal/models"
)

type GameRepository struct {
	db *sql.DB
}

func NewGameRepository(db *sql.DB) *GameRepository {
	return &GameRepository{db: db}
}

func (r *GameRepository) GetAll() ([]models.Game, error) {
	rows, err := r.db.Query("SELECT  id, title, description, release_date, developer,publisher,cover_url,platforms FROM games")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []models.Game
	for rows.Next() {
		var g models.Game
		if err := rows.Scan(&g.ID, &g.Title, &g.Description, &g.ReleaseDate, &g.Developer, &g.Publisher, &g.CoverURL, &g.Platforms); err != nil {
			return nil, err
		}
		games = append(games, g)
	}
	return games, nil
}

func (r *GameRepository) GetGameByID(id int) (*models.Game, error) {
	row := r.db.QueryRow("SELECT  id, title, description, release_date, developer,publisher,cover_url,platforms FROM games WHERE id=$1", id)

	var g models.Game
	err := row.Scan(&g.ID, &g.Title, &g.Description, &g.ReleaseDate, &g.Developer, &g.Publisher, &g.CoverURL, &g.Platforms)
	if err != nil {
		return nil, err
	}

	return &g, nil
}

func (r *GameRepository) SearchGames(query string) ([]models.Game, error) {
	rows, err := r.db.Query(`
        SELECT id, title, description, release_date 
        FROM games 
        WHERE LOWER(title) LIKE '%' || LOWER($1) || '%' 
        ORDER BY title
    `, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []models.Game
	for rows.Next() {
		var g models.Game
		err := rows.Scan(&g.ID, &g.Title, &g.Description, &g.ReleaseDate)
		if err != nil {
			return nil, err
		}
		games = append(games, g)
	}

	return games, nil
}
func (r *GameRepository) GetGenresByGameID(gameID int) ([]string, error) {
	rows, err := r.db.Query(`
        SELECT g.name
        FROM genres g
        JOIN game_genres gg ON g.id = gg.genre_id
        WHERE gg.game_id = $1
    `, gameID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genres []string
	for rows.Next() {
		var genre string
		if err := rows.Scan(&genre); err != nil {
			return nil, err
		}
		genres = append(genres, genre)
	}

	return genres, nil
}

func (r *GameRepository) GetSimilarGames(gameID int) ([]models.Game, error) {
	rows, err := r.db.Query(`
        SELECT DISTINCT g.id, g.title, g.description, g.release_date,g.cover_url
        FROM games g
        JOIN game_genres gg ON g.id = gg.game_id
        WHERE gg.genre_id IN (
            SELECT genre_id FROM game_genres WHERE game_id = $1
        )
        AND g.id != $1
        LIMIT 10
    `, gameID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []models.Game
	for rows.Next() {
		var g models.Game
		if err := rows.Scan(&g.ID, &g.Title, &g.Description, &g.ReleaseDate, &g.CoverURL); err != nil {
			return nil, err
		}
		games = append(games, g)
	}

	return games, nil
}

func (r *GameRepository) GetRecommendedGamesByUser(userID int) ([]models.Game, error) {
	query := `
		SELECT g.id, g.title, g.description, g.release_date, g.developer, g.publisher, g.cover_url, g.platforms
		FROM games g
		JOIN game_genres gg ON g.id = gg.game_id
		WHERE gg.genre_id IN (
		    SELECT gg.genre_id
		    FROM favorites f
		    JOIN game_genres gg ON f.game_id = gg.game_id
		    WHERE f.user_id = $1
		    UNION
		    SELECT gg.genre_id
		    FROM ratings r
		    JOIN game_genres gg ON r.game_id = gg.game_id
		    WHERE r.user_id = $1
		)
		AND g.id NOT IN (
		    SELECT game_id FROM ratings WHERE user_id = $1
		    UNION
		    SELECT game_id FROM favorites WHERE user_id = $1
		)
		AND g.release_date <= CURRENT_DATE -- <=== Условие: игра уже вышла
		GROUP BY g.id
		LIMIT 5
		
	`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []models.Game
	for rows.Next() {
		var g models.Game
		err := rows.Scan(&g.ID, &g.Title, &g.Description, &g.ReleaseDate, &g.Developer, &g.Publisher, &g.CoverURL, &g.Platforms)
		if err != nil {
			return nil, err
		}
		games = append(games, g)
	}

	return games, nil
}

func (r *GameRepository) GetTopRatedGames(limit int) ([]models.Game, error) {
	query := `
		SELECT g.id, g.title, g.description, g.release_date, g.developer, g.publisher, g.cover_url, g.platforms,
			   AVG(r.score) AS average_score
		FROM games g
		JOIN ratings r ON g.id = r.game_id
		GROUP BY g.id
		ORDER BY average_score DESC
		LIMIT $1;
	`
	rows, err := r.db.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []models.Game
	for rows.Next() {
		var g models.Game
		var avgScore float64
		if err := rows.Scan(&g.ID, &g.Title, &g.Description, &g.ReleaseDate, &g.Developer, &g.Publisher, &g.CoverURL, &g.Platforms, &avgScore); err != nil {
			return nil, err
		}
		g.AverageScore = &avgScore
		games = append(games, g)
	}
	return games, nil
}

func (r *GameRepository) GetRecentGames(limit int) ([]models.Game, error) {
	query := `
		SELECT id, title, description, release_date, developer, publisher, cover_url, platforms
		FROM games
		WHERE release_date <= CURRENT_DATE
		ORDER BY release_date DESC
		LIMIT $1;
	`
	rows, err := r.db.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []models.Game
	for rows.Next() {
		var g models.Game
		if err := rows.Scan(&g.ID, &g.Title, &g.Description, &g.ReleaseDate, &g.Developer, &g.Publisher, &g.CoverURL, &g.Platforms); err != nil {
			return nil, err
		}
		games = append(games, g)
	}
	return games, nil
}

func (r *GameRepository) GetUpcomingGames(limit int) ([]models.Game, error) {
	query := `
		SELECT id, title, description, release_date, developer, publisher, cover_url, platforms
		FROM games
		WHERE release_date > CURRENT_DATE
		ORDER BY release_date ASC
		LIMIT $1;
	`
	rows, err := r.db.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []models.Game
	for rows.Next() {
		var g models.Game
		if err := rows.Scan(&g.ID, &g.Title, &g.Description, &g.ReleaseDate, &g.Developer, &g.Publisher, &g.CoverURL, &g.Platforms); err != nil {
			return nil, err
		}
		games = append(games, g)
	}
	return games, nil
}

func (r *GameRepository) GetPopularGames(limit int) ([]models.Game, error) {
	query := `
		SELECT 
			g.id, g.title, g.description, g.release_date, 
			g.developer, g.publisher, g.cover_url, g.platforms,
			COUNT(DISTINCT r.id) AS recent_rating_count,
			COUNT(DISTINCT f.user_id) AS favorite_count
		FROM games g
		LEFT JOIN ratings r ON g.id = r.game_id AND r.created_at >= NOW() - INTERVAL '90 days'
		LEFT JOIN favorites f ON g.id = f.game_id
		GROUP BY g.id
		ORDER BY (COUNT(DISTINCT r.id) + COUNT(DISTINCT f.user_id)) DESC
		LIMIT $1;
	`

	rows, err := r.db.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []models.Game
	for rows.Next() {
		var g models.Game
		var recentRatings, favorites int
		if err := rows.Scan(
			&g.ID, &g.Title, &g.Description, &g.ReleaseDate,
			&g.Developer, &g.Publisher, &g.CoverURL, &g.Platforms,
			&recentRatings, &favorites,
		); err != nil {
			return nil, err
		}
		games = append(games, g)
	}
	return games, nil
}
