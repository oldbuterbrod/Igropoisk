package handlers

import (
	"igropoisk_backend/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GameHandler struct {
	service       *services.GameService
	ratingService *services.RatingService
}

func NewGameHandler(service *services.GameService) *GameHandler {
	return &GameHandler{service: service}
}

func (h *GameHandler) GetGames(c *gin.Context) {
	games, err := h.service.GetGames()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, games)
}

func (h *GameHandler) GetGameByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid game id"})
		return
	}

	game, err := h.service.GetGameByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "game not found"})
		return
	}

	c.JSON(http.StatusOK, game)
}

func (h *GameHandler) SearchGames(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "query parameter is required"})
		return
	}

	games, err := h.service.SearchGames(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to search games"})
		return
	}

	c.JSON(http.StatusOK, games)
}

func (h *GameHandler) GetGenresByGameID(c *gin.Context) {
	idStr := c.Param("id")
	gameID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid game id"})
		return
	}

	genres, err := h.service.GetGenresByGameID(gameID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get genres"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"genres": genres})
}

func (h *GameHandler) GetSimilarGames(c *gin.Context) {
	idStr := c.Param("id")
	gameID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid game id"})
		return
	}

	games, err := h.service.GetSimilarGames(gameID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get similar games"})
		return
	}

	c.JSON(http.StatusOK, games)
}

func (h *GameHandler) RateGame(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	gameID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid game id"})
		return
	}

	var input struct {
		Score int `json:"score"`
	}
	if err := c.ShouldBindJSON(&input); err != nil || input.Score < 1 || input.Score > 10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "score must be between 1 and 10"})
		return
	}

	err = h.ratingService.RateGame(userID.(int), gameID, input.Score)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to rate"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "rated"})
}

func (h *GameHandler) GetRating(c *gin.Context) {
	userID, _ := c.Get("userID") 

	gameID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid game id"})
		return
	}

	avg, err := h.ratingService.GetAverageRating(gameID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch average"})
		return
	}

	var userScore *int
	if uid, ok := userID.(int); ok {
		score, err := h.ratingService.GetUserRating(uid, gameID)
		if err == nil {
			userScore = &score
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"average":    avg,
		"user_score": userScore,
	})
}

func (h *GameHandler) SetRatingService(service *services.RatingService) {
	h.ratingService = service
}

func (h *GameHandler) GetRecommendedGames(c *gin.Context) {
	userIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID := userIDVal.(int)

	games, err := h.service.GetRecommendedGames(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch recommendations"})
		return
	}

	c.JSON(http.StatusOK, games)
}

func (h *GameHandler) GetTopRatedGames(c *gin.Context) {
	games, err := h.service.GetTopRatedGames()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch top rated games"})
		return
	}
	c.JSON(http.StatusOK, games)
}

func (h *GameHandler) GetUpcomingGames(c *gin.Context) {
	games, err := h.service.GetUpcomingGames()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch upcoming games"})
		return
	}
	c.JSON(http.StatusOK, games)
}

func (h *GameHandler) GetRecentGames(c *gin.Context) {
	games, err := h.service.GetRecentGames()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch recent games"})
		return
	}
	c.JSON(http.StatusOK, games)
}

func (h *GameHandler) GetPopularGames(c *gin.Context) {
	games, err := h.service.GetPopularGames()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch recent games"})
		return
	}
	c.JSON(http.StatusOK, games)
}
