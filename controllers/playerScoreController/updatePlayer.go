package controllers

import (
	"cicada/web-service-gin/models"
	repo "cicada/web-service-gin/repositories"
	docs "cicada/web-service-gin/req-res-models"
	"cicada/web-service-gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type NewPlayersScore struct {
	Gameid string
	Name   string
	Score  *int
}

// UpdatePlayer updates player score by ID.
// @Summary Update player score by ID
// @Description Update player score based on the provided ID.
// @Tags Player Scores
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Security bearer
// @Param Authorization header string true "Bearer "
// @Param requestBody body docs.PlayerScore true "Request Body"
// @Success 200 {object} docs.PlayerScore
// @Failure 400 {object} docs.ErrorResponse
// @Failure 401 {object} docs.ErrorResponse
// @Failure 500 {object} docs.ErrorResponse
// @Router /player/update-player-score/{id} [put]
func UpdatePlayer(serverConfigs utils.ServerConfigs) func(*gin.Context) {
	return func(c *gin.Context) {

		logger := log.Logger.With().Str("func", "UpdatePlayer").Logger()
		id := c.Param("id")
		var requestBody NewPlayersScore

		err := c.BindJSON(&requestBody)
		if err != nil {
			logger.Error().Err(err).Msg("Invalid request payload")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		existingPlayerScore, err := repo.GetPlayerScoreByID(serverConfigs.DB, id)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to get playerScore")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get playerScore"})
			return
		}
		if requestBody.Gameid != "" {
			existingPlayerScore.Gameid = requestBody.Gameid
		}
		if requestBody.Score != nil {
			existingPlayerScore.Score = *requestBody.Score
		}
		if requestBody.Name != "" {
			existingPlayerScore.Name = requestBody.Name
		}

		updatedPlayerScore, err := repo.UpdatePlayerScore(serverConfigs.DB, id, existingPlayerScore)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to update playerScore")
			errorResponse := docs.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Failed to update playerScore",
			}
			c.JSON(http.StatusInternalServerError, errorResponse)
			return
		}

		c.JSON(http.StatusOK, gin.H{

			"updatedPlayerScore": updatedPlayerScore,
		})
	}
}

// UpdatePlayer updates players score by gameid.
// @Summary Update PlayersScore by gameid
// @Description Update player scores based on the provided gameid.
// @Tags Player Scores
// @Accept json
// @Produce json
// @Param gameid path string true "gameid"
// @Security bearer
// @Param Authorization header string true "Bearer "
// @Param requestBody body docs.PlayerScore true "Request Body"
// @Success 200 {object} docs.PlayersScore
// @Failure 400 {object} docs.ErrorResponse
// @Failure 401 {object} docs.ErrorResponse
// @Failure 500 {object} docs.ErrorResponse
// @Router /player/update-player/{gameid} [put]
func UpdateAllPlayerGameid(serverConfigs utils.ServerConfigs) func(*gin.Context) {
	return func(c *gin.Context) {

		logger := log.Logger.With().Str("func", "UpdatePlayer").Logger()
		gameid := c.Param("gameid")
		var requestBody models.PlayerScore

		err := c.BindJSON(&requestBody)
		if err != nil {
			logger.Error().Err(err).Msg("Invalid request payload")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		updatedPlayerScore, err := repo.UpdatePlayerScoreGameId(serverConfigs.DB, gameid, requestBody.Gameid)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to update playerScore")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update playerScore"})
			return
		}

		c.JSON(http.StatusOK, gin.H{

			"updatedPlayerScores": updatedPlayerScore,
		})
	}
}
