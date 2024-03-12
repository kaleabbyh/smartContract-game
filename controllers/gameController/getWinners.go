package controller

import (
	//"cicada/web-service-gin/models"
	"cicada/web-service-gin/models"
	repo "cicada/web-service-gin/repositories"
	docs "cicada/web-service-gin/req-res-models"
	"cicada/web-service-gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// GetWinner get winners with their won amount  .
// @Summary get winners with their won amount  by gameid
// @Description get winners with their won amount based on the provided gameid.
// @Tags Play
// @Accept json
// @Produce json
// @Param gameid path string true "gameid"
// @Security bearer
// @Param Authorization header string true "Bearer "
// @Success 200 {object} docs.AllWinners
// @Failure 401 {object} docs.ErrorResponse
// @Failure 500 {object} docs.ErrorResponse
// @Router /play/get-winners/{gameid} [get]
func GetWinner(serverConfigs utils.ServerConfigs) func(*gin.Context) {
	return func(c *gin.Context) {

		logger := log.Logger.With().Str("func", "GetWinner").Logger()
		gameid := c.Param("gameid")

		winnersForGame, err := repo.GetWinnerForGame(serverConfigs.DB, gameid)
		if err != nil {
			logger.Error().Err(err).Msgf("Error: %s", err.Error())
			errorResponse := docs.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error: %s" + err.Error(),
			}

			c.JSON(http.StatusInternalServerError, errorResponse)
			return
		}

		PlayerScores := winnersForGame.PlayerScore
		pools := winnersForGame.Players

		highestScore := GetHighestScore(PlayerScores)
		winners := GetWinnersList(PlayerScores, highestScore)

		var AllWinners []Winner
		var winnerNames []string
		for _, winner := range winners {
			winnerNames = append(winnerNames, winner.Name)

		}

		// Create a channel for collecting results from goroutines
		resultCh := make(chan Winner)

		// goroutines to process winners concurrently
		for _, winner := range winners {
			go func(winner models.PlayerScore) {
				var won []string
				var sharedAmount float32 = float32(0)
				var notSharedAmount float32 = float32(0)

				for _, pool := range pools {
					if IsContains(winner.Name, pool.Owners) {
						winnerCount := CountWinnersPerSinglePool(winnerNames, pool.Owners)
						if winnerCount > 1 {
							won = append(won, pool.ID)
							sharedAmount += (pool.Value / float32(winnerCount))
						} else {
							won = append(won, pool.ID)
							notSharedAmount += pool.Value
						}
					}
				}

				resultCh <- Winner{
					ID:              winner.ID,
					PoolsWonID:      won,
					Name:            winner.Name,
					NotsharedAmount: notSharedAmount,
					SharedAmount:    sharedAmount,
					TotalAmount:     sharedAmount + notSharedAmount,
				}
			}(winner)
		}

		// Collect results from goroutines
		for range winners {
			winner := <-resultCh
			AllWinners = append(AllWinners, winner)
		}

		c.JSON(200, gin.H{
			"Winners": AllWinners,
		})
	}
}
