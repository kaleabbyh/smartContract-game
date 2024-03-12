package controllers

import (
	"cicada/web-service-gin/models"
	repo "cicada/web-service-gin/repositories"
	docs "cicada/web-service-gin/req-res-models"
	utils "cicada/web-service-gin/utils"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// GetAllPools fetches all pools.
// @Summary Get all pools with their poolis.
// @Description Retrieves a list of all pools.
// @Tags Pools
// @Accept json
// @Produce json
// @Security bearer
// @Param Authorization header string true "Bearer"
// @Success 200 {array} docs.Pool
// @Failure 401 {object} docs.ErrorResponse
// @Failure 500 {object} docs.ErrorResponse
// @Router /pool/get-all-pools [get]
func GetAllPools(serverConfigs utils.ServerConfigs) func(*gin.Context) {
	return func(c *gin.Context) {
		logger := log.Logger.With().Str("func", "GetAllPools").Logger()

		poolsChan := make(chan []models.Pool)
		errChan := make(chan error)

		go func() {
			pools, err := repo.GetAllPools(serverConfigs.DB)
			if err != nil {
				errChan <- err
				return
			}
			poolsChan <- pools
		}()

		select {
		case pools := <-poolsChan:
			c.JSON(http.StatusOK, gin.H{
				"status": "success",
				"pools":  pools,
			})
		case err := <-errChan:
			logger.Error().Err(err).Msgf("Failed to get pools: %s", err)
			errorResponse := docs.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Failed to get pools",
			}
			c.JSON(http.StatusInternalServerError, errorResponse)
		}
	}
}
