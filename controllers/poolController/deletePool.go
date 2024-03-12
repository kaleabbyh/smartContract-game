package controllers

import (
	repo "cicada/web-service-gin/repositories"
	docs "cicada/web-service-gin/req-res-models"
	"cicada/web-service-gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// DeletePoolWithPoollisByID Delete pool with respective poollis .
// @Summary Delete pool with poollis by pool ID
// @Description Delete pool with poollis  based on the provided pool ID.
// @Tags Pools
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Security bearer
// @Param Authorization header string true "Bearer "
// @Success 200 {object} docs.ErrorResponse
// @Failure 401 {object} docs.ErrorResponse
// @Failure 500 {object} docs.ErrorResponse
// @Router /pool/delete-pool/{id} [delete]
func DeletePoolWithPoollisByID(serverConfig utils.ServerConfigs) func(*gin.Context) {
	return func(c *gin.Context) {

		logger := log.Logger.With().Str("func", "DeletePoolWithPoollisByID").Logger()
		poolID := c.Param("id")

		err := repo.DeletePoolWithPoollisByID(serverConfig.DB, poolID)
		if err != nil {
			logger.Error().Err(err).Msgf("Failed to delete pool: %s", err)
			errorResponse := docs.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Failed to delete pool: " + poolID,
			}
			c.JSON(http.StatusInternalServerError, errorResponse)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"Mesaage": "successfully deleted pool",
		})
	}
}
