package main

import (
	"bytes"
	"cicada/web-service-gin/config"
	player "cicada/web-service-gin/controllers/playerScoreController"
	pool "cicada/web-service-gin/controllers/poolController"

	"cicada/web-service-gin/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	adminToken string
	userToken  string
	client     *http.Client
	logger     zerolog.Logger
)

func init() {
	config.LoadConfig()
	adminToken = os.Getenv("ADMIN_TOKEN")
	userToken = os.Getenv("USER_TOKEN")
	client = &http.Client{}
	logger = log.Logger.With().Str("func", "test").Logger()
}

func main() {

	config.Migrate()
	logger.Info().Msg("Successfully migrated")

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	//create pool with its poollis
	poolRequest := pool.Players{
		Players: []pool.Player{
			{Name: "Player 1", Amount: 20},
			{Name: "Player 2", Amount: 20},
			{Name: "Player 3", Amount: 40},
			{Name: "Player 4", Amount: 80},
			{Name: "Player 5", Amount: 100},
		},
	}
	poolURL := "http://localhost:8080/pool/create-pool"
	poolResponse, err := makeRequest("POST", poolURL, poolRequest, userToken)
	if err != nil || poolResponse == nil {
		logError("Failed to create pool", err)
		return
	}

	gameID, ok := poolResponse["Gameid"].(string)
	if !ok {
		logger.Fatal().Msg("Failed to extract Gameid from the response")
		return
	}
	PoolID, ok := poolResponse["PoolID"].(string)
	if !ok {
		logger.Fatal().Msg("Failed to extract PoolID from the response")
		return
	}

	logger.Info().Msg("Pool successfully created")
	logger.Info().Msgf("GameID: %s", gameID)

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	//create playeScores
	playersScore := player.PlayersScore{
		Gameid: gameID,
		PlayersScore: []models.PlayerScore{
			{Name: "Player 1", Score: 10},
			{Name: "Player 2", Score: 15},
			{Name: "Player 3", Score: 15},
			{Name: "Player 4", Score: 15},
			{Name: "Player 5", Score: 15},
		},
	}
	playerURL := "http://localhost:8080/player/create-player-score"
	playerResponse, err := makeRequest("POST", playerURL, playersScore, adminToken)
	if err != nil || playerResponse == nil {
		logError("Failed to create player scores", err)
		return
	} else {
		logger.Info().Msg("Player scores successfully created")
	}

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	//winners of the game
	getWinnersURL := "http://localhost:8080/play/get-winners/" + gameID
	getWinnersResponse, err := makeRequest("GET", getWinnersURL, nil, userToken)
	if err != nil {
		logError("Failed to get winners", err)
		return
	} else {
		fmt.Print("\n<<<<<<<<<<<<<<<<<<<<<<<< WINNERS FOR THE GAME >>>>>>>>>>>>>>>>>>>>>>>>>>")
		jsonData, err := json.MarshalIndent(getWinnersResponse, "", "    ")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("\n%s\n\n", jsonData)
	}

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	// update the pool or game state to completed
	newPool := models.Pool{
		GameStatus: "completed",
	}
	updatePoolURL := "http://localhost:8080/pool/update-pool/" + PoolID
	updatePoolResponse, err := makeRequest("PUT", updatePoolURL, newPool, adminToken)
	if err != nil || updatePoolResponse == nil {
		logError("Failed to update pool", err)
		return
	} else {
		logger.Info().Msg("Pool successfully updated to completed")
	}

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	//delte pool of the game
	deletePoolURL := "http://localhost:8080/pool/delete-pool/" + PoolID
	_, err = makeRequest("DELETE", deletePoolURL, nil, adminToken)
	if err != nil {
		logError("Failed to delete pool", err)
		return
	} else {
		logger.Info().Msg("Pool successfully deleted")
	}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
// Function to make requests to different methods
func makeRequest(method, url string, requestData interface{}, authToken string) (map[string]interface{}, error) {
	payload, err := json.Marshal(requestData)
	if err != nil {
		logError("Failed to marshal request", err)
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		logError("Failed to create request", err)
		return nil, err
	}

	req.Header.Set("Authorization", authToken)
	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		logError("Failed to send request", err)
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		err := fmt.Errorf("request failed with status code %d", response.StatusCode)
		return nil, err
	}

	var result map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		logError("Failed to decode response", err)
		return nil, err
	}

	return result, nil
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
// Error logging functions using zerolog
func logError(message string, err error) {
	logger.Error().Err(err).Msgf("%s: %s", message, err.Error())
}
