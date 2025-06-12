package repositories

import (
	"database/sql"
	"errors"
	"igropoisk_backend/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(username, email, passwordHash string) error {
	_, err := r.db.Exec("INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)", username, email, passwordHash)
	return err
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	row := r.db.QueryRow("SELECT id, username, email, password_hash FROM users WHERE email=$1", email)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	return user, err
}

func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	user := &models.User{}
	row := r.db.QueryRow("SELECT id, username, email FROM users WHERE id=$1", id)
	err := row.Scan(&user.ID, &user.Username, &user.Email)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	return user, err
}
