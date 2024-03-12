package controllers

import (
	repo "cicada/web-service-gin/repositories"
	docs "cicada/web-service-gin/req-res-models"
	"cicada/web-service-gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// GetPoolByID get pool with respective poollis .
// @Summary get pool with poollis by pool ID
// @Description get pool with poollis  based on the provided pool ID.
// @Tags Pools
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Security bearer
// @Param Authorization header string true "Bearer "
// @Success 200 {object} docs.Pool
// @Failure 401 {object} docs.ErrorResponse
// @Failure 500 {object} docs.ErrorResponse
// @Router /pool/get-pool/{id} [get]
func GetPoolByID(serverConfigs utils.ServerConfigs) func(*gin.Context) {
	return func(c *gin.Context) {

		logger := log.Logger.With().Str("func", "GetPoolByID").Logger()

		id := c.Param("id")
		pool, err := repo.GetPoolByID(serverConfigs.DB, id)
		if err != nil {
			logger.Error().Err(err).Msgf("Failed to get pool by ID: %s", err)
			errorResponse := docs.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Failed to get pool by id: " + id,
			}
			c.JSON(http.StatusInternalServerError, errorResponse)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"pool":   pool,
		})
	}
}
