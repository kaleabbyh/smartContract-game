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

// UpdatePool updates pool by id.
// @Summary updates pool by id
// @Description updates pool based on the provided id.
// @Tags Pools
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Security bearer
// @Param Authorization header string true "Bearer "
// @Param requestBody body docs.Pool true "Request Body"
// @Success 200 {object} docs.Pool
// @Failure 400 {object} docs.ErrorResponse
// @Failure 401 {object} docs.ErrorResponse
// @Failure 500 {object} docs.ErrorResponse
// @Router /pool/update-pool/{id} [put]
func UpdatePool(serverConfigs utils.ServerConfigs) func(*gin.Context) {
	return func(c *gin.Context) {

		logger := log.Logger.With().Str("func", "UpdatePool").Logger()
		id := c.Param("id")
		var requestBody models.Pool

		err := c.BindJSON(&requestBody)
		if err != nil {
			logger.Error().Err(err).Msg("Invalid request payload")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		updatedPool, err := repo.UpdatePool(serverConfigs.DB, id, requestBody)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to update Pool")
			errorResponse := docs.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "invalid request body",
			}
			c.JSON(http.StatusInternalServerError, errorResponse)
			return
		}

		c.JSON(http.StatusOK, gin.H{

			"updatedPool": updatedPool,
		})
	}
}
