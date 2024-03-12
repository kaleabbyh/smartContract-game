package controllers

import (
	docs "cicada/web-service-gin/req-res-models"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type PoolToSort struct {
	PoolID  string `json:"poolID"`
	Pooliis []Poollis
}

type Poollis struct {
	ID     string         `json:"id"`
	PoolID string         `json:"poolID,omitempty"`
	Value  float32        `json:"value"`
	Owners pq.StringArray `json:"owners"`
}

// SortPool sort pool with number of owners .
// @Summary sort pool with number of owners
// @Description sort pool with number of owners based on the provided poolli Data.
// @Tags Pools
// @Accept json
// @Produce json
// @Security bearer
// @Param Authorization header string true "Bearer "
// @Param body body docs.PoolToSort true "Player data"
// @Success 200 {object} docs.PoolToSort
// @Failure 401 {object} docs.ErrorResponse
// @Failure 500 {object} docs.ErrorResponse
// @Router /pool/sort-pool [post]
func SortPool() func(*gin.Context) {
	return func(c *gin.Context) {

		logger := log.Logger.With().Str("func", "SortPool").Logger()

		var requestBody PoolToSort
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Error().Err(err).Msgf("invalid request body: %s", err)
			errorResponse := docs.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "invalid request body",
			}
			c.JSON(http.StatusNotAcceptable, errorResponse)
			return
		}

		// Sort the Poollis array by number of Owners in ascending order
		sort.Slice(requestBody.Pooliis, func(i, j int) bool {
			return len(requestBody.Pooliis[i].Owners) > len(requestBody.Pooliis[j].Owners)
		})

		logger.Info().Msg("pool sorted")
		c.JSON(200, gin.H{
			"Pool": requestBody,
		})
	}
}
