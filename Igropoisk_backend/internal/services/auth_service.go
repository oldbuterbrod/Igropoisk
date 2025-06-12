package services

import (
    "time"

    "github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
    secret string
}

func NewAuthService(secret string) *AuthService {
    return &AuthService{secret: secret}
}

func (s *AuthService) GenerateToken(userID int) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(24 * time.Hour).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(s.secret))
}
