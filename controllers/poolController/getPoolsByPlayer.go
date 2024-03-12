package controllers

import (
	repo "cicada/web-service-gin/repositories"
	docs "cicada/web-service-gin/req-res-models"
	"cicada/web-service-gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// GetAllPoollisPlayerIsIn get all poollis a player is in .
// @Summary get all poollis a player is in
// @Description get all poollis a player is in  based on the provided player and poolID.
// @Tags Pools
// @Accept json
// @Produce json
// @Param player path string true "player"
// @Param poolID path string true "poolID"
// @Security bearer
// @Param Authorization header string true "Bearer "
// @Success 200 {array} docs.Poolli
// @Failure 401 {object} docs.ErrorResponse
// @Failure 500 {object} docs.ErrorResponse
// @Router /pool/get-poollis/{poolID}/{player} [get]
func GetAllPoollisPlayerIsIn(serverConfigs utils.ServerConfigs) func(*gin.Context) {
	return func(c *gin.Context) {
		logger := log.Logger.With().Str("func", "GetAllPoollisPlayerIsIn").Logger()

		player := c.Param("player")
		poolID := c.Param("poolID")
		poollis, err := repo.GetPoollisplayerIsIn(serverConfigs.DB, poolID, player)
		if err != nil {
			logger.Error().Err(err).Msgf("Failed to get poollis: %s", err.Error())
			errorResponse := docs.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Failed to get poollis: " + err.Error(),
			}
			c.JSON(http.StatusInternalServerError, errorResponse)
		}
		c.JSON(http.StatusOK, gin.H{"poollis": poollis})
	}
}
