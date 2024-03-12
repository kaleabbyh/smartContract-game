package controllers

import (
	repo "cicada/web-service-gin/repositories"
	docs "cicada/web-service-gin/req-res-models"
	"cicada/web-service-gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// GetPlayerScoreByGameID fetches player scores by game ID.
// @Summary Get player scores by game ID
// @Description Retrieves player scores based on the provided game ID.
// @Tags Player Scores
// @Accept json
// @Produce json
// @Param gameid path string true "Game ID"
// @Security bearer
// @Param Authorization header string true "Bearer "
// @Success 200 {object} docs.PlayerScoreResponse
// @Failure 401 {object} docs.ErrorResponse
// @Failure 500 {object} docs.ErrorResponse
// @Router /player/get-player-scores/{gameid} [get]
func GetPlayerScoreByGameID(serverConfigs utils.ServerConfigs) func(*gin.Context) {
	return func(c *gin.Context) {

		logger := log.Logger.With().Str("func", "GetPlayerScoreByGameID").Logger()

		gameid := c.Param("gameid")
		pool, err := repo.GetPlayerScoreByGameId(serverConfigs.DB, gameid)
		if err != nil {
			logger.Error().Err(err).Msgf("Failed to get player score by .GameID: %s", err)
			errorResponse := docs.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Failed to get player score by GameID",
			}
			c.JSON(http.StatusInternalServerError, errorResponse)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":       "success",
			"playerScores": pool,
		})
	}
}
