package models

import "time"

type User struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
}

type Game struct {
	ID             int       `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	ReleaseDate    time.Time `json:"release_date"`
	Developer      string    `json:"developer"`
	Publisher      string    `json:"publisher"`
	CoverURL       string    `json:"cover_url"`
	Platforms      string    `json:"platforms"`
	Review         string    `json:"review,omitempty"`
	AverageScore   *float64  `json:"average_score,omitempty"`
	RatingCount    int       `json:"rating_count,omitempty"`    
	FavoritesCount int       `json:"favorites_count,omitempty"`
}

type Review struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	GameID    int       `json:"gameId"` 
	Text      string    `json:"text"`
	Score     int       `json:"score"`
	CreatedAt time.Time `json:"created_at"`
	User      *User     `json:"user,omitempty"`
	GameTitle string    `json:"gameTitle"`
}

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GameGenre struct {
	GameID  int `json:"game_id"`
	GenreID int `json:"genre_id"`
}

type Favorite struct {
	UserID int       `json:"user_id"`
	GameID int       `json:"game_id"`
	Added  time.Time `json:"added_at"`
}

type Rating struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	GameID    int       `json:"game_id"`
	Score     int       `json:"score"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"Title"`
}
