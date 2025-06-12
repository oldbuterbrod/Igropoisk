package main

import (
	"database/sql"
	"fmt"
	"igropoisk_backend/internal/config"
	"igropoisk_backend/internal/handlers"
	"igropoisk_backend/internal/middleware"
	"igropoisk_backend/internal/repositories"
	"igropoisk_backend/internal/services"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.LoadConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	defer db.Close()

	// Игровой сервис
	gameRepo := repositories.NewGameRepository(db)
	gameService := services.NewGameService(gameRepo)
	gameHandler := handlers.NewGameHandler(gameService)

	// Юзер сервисы для авторизации
	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(cfg.JWTSecret)
	authHandler := handlers.NewAuthHandler(authService, userRepo)

	// Репозиторий и сервис рейтинга
	ratingRepo := repositories.NewRatingRepository(db)
	ratingService := services.NewRatingService(ratingRepo)
	ratingHandler := handlers.NewRatingHandler(ratingService)
	// Подключаем сервис рейтинга к gameHandler
	gameHandler.SetRatingService(ratingService)

	reviewRepo := repositories.NewReviewRepository(db)
	reviewService := services.NewReviewService(reviewRepo)
	reviewHandler := handlers.NewReviewHandler(reviewService)

	

	favoriteRepo := repositories.NewFavoriteRepository(db)
	favoriteService := services.NewFavoriteService(favoriteRepo)
	favoriteHandler := handlers.NewFavoriteHandler(favoriteService)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	api := r.Group("/api")
	{
		// Игровые роуты
		api.GET("/games", gameHandler.GetGames)
		api.GET("/games/:id", gameHandler.GetGameByID)
		api.GET("/games/search", gameHandler.SearchGames)
		api.GET("/games/:id/genres", gameHandler.GetGenresByGameID)
		api.GET("/games/:id/similar", gameHandler.GetSimilarGames)

		// Группа авторизации
		authGroup := api.Group("/auth")
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
		authGroup.GET("/me", middleware.JWTMiddleware(cfg.JWTSecret), authHandler.GetMe)

		// Оценки (рейтинг)
		api.POST("/games/:id/rate", middleware.JWTMiddleware(cfg.JWTSecret), gameHandler.RateGame)
		api.GET("/games/:id/rating", middleware.OptionalJWTMiddleware(cfg.JWTSecret), gameHandler.GetRating)
		api.GET("/user/ratings", middleware.JWTMiddleware(cfg.JWTSecret), ratingHandler.GetUserRatings)
		api.POST("/games/:id/reviews", middleware.JWTMiddleware(cfg.JWTSecret), reviewHandler.CreateReview)
		api.GET("/games/:id/reviews", reviewHandler.GetReviews)
		api.GET("/user/reviews", middleware.JWTMiddleware(cfg.JWTSecret), reviewHandler.GetUserReviews)
		api.POST("/games/:id/favorite", middleware.JWTMiddleware(cfg.JWTSecret), favoriteHandler.AddFavorite)
		api.DELETE("/games/:id/favorite", middleware.JWTMiddleware(cfg.JWTSecret), favoriteHandler.RemoveFavorite)
		api.GET("/user/favorites", middleware.JWTMiddleware(cfg.JWTSecret), favoriteHandler.GetFavorites)
		api.GET("/games/recommendations", middleware.JWTMiddleware(cfg.JWTSecret), gameHandler.GetRecommendedGames)
		api.GET("/games/top", gameHandler.GetTopRatedGames)
		api.GET("/games/recent", gameHandler.GetRecentGames)
		api.GET("/games/upcoming", gameHandler.GetUpcomingGames)
		api.GET("/games/popular", gameHandler.GetPopularGames)
		api.DELETE("/games/:id/review", middleware.JWTMiddleware(cfg.JWTSecret), reviewHandler.DeleteReview)
		api.DELETE("/games/:id/rating", middleware.JWTMiddleware(cfg.JWTSecret), ratingHandler.DeleteRating)

	}

	r.Run(":" + cfg.ServerPort)
}
