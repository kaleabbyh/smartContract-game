package controllers

import (
	"cicada/web-service-gin/models"
	repo "cicada/web-service-gin/repositories"
	docs "cicada/web-service-gin/req-res-models"
	"cicada/web-service-gin/utils"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type PlayersScore struct {
	Gameid       string
	PlayersScore []models.PlayerScore
}

// CreatePlayer creates players.
// @Summary Registers players with their scores.
// @Description Registers players with their scores.
// @Tags Player Scores
// @Accept json
// @Produce json
// @Security bearer
// @Param Authorization header string true "Bearer"
// @Param body body docs.PlayersScore true "Player score data"
// @Success 200 {object} docs.PlayerScoreResponse
// @Failure 401 {object} docs.ErrorResponse
// @Failure 500 {object} docs.ErrorResponse
// @Router /player/create-player-score [post]
func CreatePlayer(serverConfigs utils.ServerConfigs) func(*gin.Context) {
	return func(c *gin.Context) {
		logger := log.Logger.With().Str("func", "CreatePlayer").Logger()
		var requestBody PlayersScore

		err := c.BindJSON(&requestBody)
		if err != nil {
			logger.Error().Err(err).Msg("Invalid request payload")
			errorResponse := docs.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Invalid request payload",
			}
			c.JSON(http.StatusBadRequest, errorResponse)
			return
		}

		var playerScoreResponse []models.PlayerScore
		var Gameid string = requestBody.Gameid

		var wg sync.WaitGroup
		playerScoreCh := make(chan models.PlayerScore, len(requestBody.PlayersScore))

		for _, playerScore := range requestBody.PlayersScore {
			wg.Add(1)
			go func(playerScore models.PlayerScore) {
				defer wg.Done()

				newPlayer := models.PlayerScore{
					Gameid: Gameid,
					Name:   playerScore.Name,
					Score:  playerScore.Score,
				}

				SinglePlayerScore, err := repo.CreateNePlayer(serverConfigs.DB, newPlayer)
				if err != nil {
					logger.Error().Err(err).Msg("Failed to create playerScore")
					errorResponse := docs.ErrorResponse{
						Code:    http.StatusInternalServerError,
						Message: "Failed to create playerScore",
					}
					c.JSON(http.StatusBadRequest, errorResponse)
					return
				}

				playerScoreCh <- SinglePlayerScore
			}(playerScore)
		}

		go func() {
			wg.Wait()
			close(playerScoreCh)
		}()

		for playerScore := range playerScoreCh {
			playerScoreResponse = append(playerScoreResponse, playerScore)
		}

		c.JSON(http.StatusOK, gin.H{
			"playerScore": playerScoreResponse,
		})
	}
}
