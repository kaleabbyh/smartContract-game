package controllers

import (
	"cicada/web-service-gin/models"
	repo "cicada/web-service-gin/repositories"
	docs "cicada/web-service-gin/req-res-models"

	"cicada/web-service-gin/utils"

	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type Player struct {
	Name   string
	Amount float32
}

type Players struct {
	Players []Player
}

type PoolResponse struct {
	ID     string
	Value  float32
	Owners pq.StringArray
}

// CreatePool creates pool.
// @Summary Registers pool with their poollis.
// @Description Registers pool with their poollis.
// @Tags Pools
// @Accept json
// @Produce json
// @Security bearer
// @Param Authorization header string true "Bearer "
// @Param body body docs.Players true "Player data"
// @Success 200 {object} docs.CreatePoolResponse
// @Failure 401 {object} docs.ErrorResponse
// @Failure 500 {object} docs.ErrorResponse
// @Router /pool/create-pool [post]
func CreatePool(serverConfigs utils.ServerConfigs) func(*gin.Context) {
	return func(c *gin.Context) {

		logger := log.Logger.With().Str("func", "CreatePool").Logger()
		var requestBody Players
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Error().Err(err).Msgf("invalid request body: %s", err)
			c.JSON(http.StatusNotAcceptable, err)
			return
		}

		players := requestBody.Players
		sort.Slice(players, func(i, j int) bool {
			return players[i].Amount < players[j].Amount
		})

		minValue := players[0].Amount
		count := 0

		var pools []PoolResponse

		if len(players) <= 0 {
			logger.Error().Msg("No player found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No player found"})
			return
		}

		gameid, _ := utils.GenerateUUID()

		NewPoolRequest := &models.Pool{
			Gameid: gameid,
		}

		newpoolCreated, err := repo.CreateNewPool(serverConfigs.DB, NewPoolRequest)
		if err != nil {
			logger.Error().Err(err).Msgf("Failed to create pool: %s", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create pool"})
			return
		}

		for len(players) > 0 {
			for i := range players {
				players[i].Amount -= minValue
			}

			var owners []string
			for _, player := range players {
				owners = append(owners, player.Name)
			}

			newPoolliPayload := &models.Poolli{
				PoolID: newpoolCreated.ID,
				Value:  minValue * float32(len(players)),
				Owners: owners,
			}

			newpoolliCreated, err := repo.CreateNewPoolli(serverConfigs.DB, newPoolliPayload)
			if err != nil {
				logger.Error().Err(err).Msgf("Failed to create poolli: %s", err)
				errorResponse := docs.ErrorResponse{
					Code:    http.StatusInternalServerError,
					Message: "Failed to create pool-li",
				}
				c.JSON(http.StatusInternalServerError, errorResponse)
				return
			}

			pool := PoolResponse{
				ID:     newpoolliCreated.ID,
				Value:  minValue * float32(len(players)),
				Owners: owners,
			}

			var updatedArray []Player
			for _, player := range players {
				if player.Amount > 0 {
					updatedArray = append(updatedArray, player)
				}
			}
			players = updatedArray

			if len(players) > 0 {
				minValue = players[0].Amount
			}

			pools = append(pools, pool)
			count++
		}

		c.JSON(200, gin.H{
			"Gameid":  newpoolCreated.Gameid,
			"PoolID":  newpoolCreated.ID,
			"Poollis": pools,
		})
	}
}
