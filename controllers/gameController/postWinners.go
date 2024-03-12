package controller

import (
	"cicada/web-service-gin/models"
	docs "cicada/web-service-gin/req-res-models"
	"net/http"
	"sort"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type PlayerScore struct {
	ID    string
	Name  string
	Score int
}

type Players struct {
	ID     string
	Name   string
	Amount float32
}
type GameRequest struct {
	PlayerScore []models.PlayerScore
	Players     []Players
}

type SortedPoolResponse struct {
	ID     string
	Count  int
	Value  float32
	Owners pq.StringArray
}

type Winner struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	NotsharedAmount float32  `json:"not_shared_amount"`
	SharedAmount    float32  `json:"shared_amount"`
	TotalAmount     float32  `json:"total_amount"`
	PoolsWonID      []string `json:"pools_won_id"`
}

// PostWinners identify winners from  the given player using thier scores.
// @Summary identify winners from  the given player using thier scores.
// @Description identify winners from  the given player using thier scores.
// @Tags Play
// @Accept json
// @Produce json
// @Security bearer
// @Param Authorization header string true "Bearer"
// @Param body body docs.GameRequest true "Player data"
// @Success 200 {object} docs.AllWinners
// @Failure 401 {object} docs.ErrorResponse
// @Failure 500 {object} docs.ErrorResponse
// @Router /play/winners [post]
func PostWinners() func(*gin.Context) {
	return func(c *gin.Context) {
		logger := log.Logger.With().Str("func", "PostWinners").Logger()

		var requestBody GameRequest
		if err := c.BindJSON(&requestBody); err != nil {
			logger.Error().Err(err).Msgf("invalid request body: %s", err)
			errorResponse := docs.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "invalid request body: %s" + err.Error(),
			}
			c.JSON(http.StatusInternalServerError, errorResponse)
			return
		}

		PlayerScores := requestBody.PlayerScore
		pools := CreatePool(requestBody.Players)
		highestScore := GetHighestScore(PlayerScores)
		winners := GetWinnersList(PlayerScores, highestScore)

		var AllWinners []Winner
		var winnerNames []string
		for _, winner := range winners {
			winnerNames = append(winnerNames, winner.Name)
		}

		// Create a channel for collecting results from goroutines
		resultCh := make(chan Winner)

		//goroutines to process winners concurrently
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

func CreatePool(players []Players) []SortedPoolResponse {

	logger := log.Logger.With().Str("func", "CreatePool").Logger()

	sort.Slice(players, func(i, j int) bool {
		return players[i].Amount < players[j].Amount
	})

	minValue := players[0].Amount
	count := 1

	var poollis []SortedPoolResponse

	if len(players) <= 0 {
		logger.Error().Msg("No player found")
		return nil
	}

	for len(players) > 0 {
		for i := range players {
			players[i].Amount -= minValue
		}

		var owners []string
		for _, player := range players {
			owners = append(owners, player.Name)
		}

		pool := SortedPoolResponse{
			ID:     strconv.Itoa(count),
			Count:  count,
			Value:  minValue * float32(len(players)),
			Owners: owners,
		}

		var updatedArray []Players
		for _, player := range players {
			if player.Amount > 0 {
				updatedArray = append(updatedArray, player)
			}
		}
		players = updatedArray

		if len(players) > 0 {
			minValue = players[0].Amount
		}

		poollis = append(poollis, pool)
		count++
	}

	return poollis
}

func GetHighestScore(PlayerScores []models.PlayerScore) int {
	highestScore := PlayerScores[0].Score
	for _, score := range PlayerScores {
		if score.Score > highestScore {
			highestScore = score.Score
		}
	}
	return highestScore
}

func GetWinnersList(PlayerScores []models.PlayerScore, highestScore int) []models.PlayerScore {
	winnersCh := make(chan models.PlayerScore, len(PlayerScores))
	var wg sync.WaitGroup

	for _, player := range PlayerScores {
		player := player
		wg.Add(1)
		go func() {
			defer wg.Done()
			if player.Score == highestScore {
				winnersCh <- player
			}
		}()
	}

	go func() {
		wg.Wait()
		close(winnersCh)
	}()

	var winners []models.PlayerScore
	for winner := range winnersCh {
		winners = append(winners, winner)
	}

	return winners
}

func CountWinnersPerSinglePool(winners []string, owners []string) int {
	count := 0
	for _, winner := range winners {
		for _, owner := range owners {
			if winner == owner {
				count++
				break
			}
		}
	}
	return count
}

func IsContains(str string, words []string) bool {
	for _, word := range words {
		if word == str {
			return true
		}
	}
	return false
}
