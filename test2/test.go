package main

import (
	"bytes"
	"cicada/web-service-gin/config"
	pool "cicada/web-service-gin/controllers/poolController"

	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	AdminToken string
	UserToken  string
	client     *http.Client
	logger     zerolog.Logger
)

func init() {
	config.LoadConfig()
	AdminToken = os.Getenv("ADMIN_TOKEN")
	UserToken = os.Getenv("USER_TOKEN")
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
	poolResponse, err := makeRequest("POST", poolURL, poolRequest, UserToken)
	if err != nil || poolResponse == nil {
		logError("Failed to create pool", err)
		return
	}

	PoolID, ok := poolResponse["PoolID"].(string)
	if !ok {
		logger.Fatal().Msg("Failed to extract PoolID from the response")
		return
	}
	logger.Info().Msg("Pool successfully created")

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	//GET POOL BY ID WITH ALL ITS POOLLIS
	getWinnersURL := "http://localhost:8080/pool/get-pool/" + PoolID
	getWinnersResponse, err := makeRequest("GET", getWinnersURL, nil, UserToken)
	if err != nil {
		logError("Failed to get pool by ID", err)
		return
	} else {
		fmt.Print("\n<<<<<<<<<<<<<<<<<<<<<<<< ALL POOLLIS OF POOLID = " + PoolID + " >>>>>>>>>>>>>>>>>>>>>>>>>>")
		jsonData, err := json.MarshalIndent(getWinnersResponse, "", "    ")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("\n%s\n\n", jsonData)
	}

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	//pools each player is in
	for _, player := range poolRequest.Players {

		getWinnersURL := "http://localhost:8080/pool/get-poollis/" + PoolID + "/" + player.Name
		getWinnersResponse, err := makeRequest("GET", getWinnersURL, nil, UserToken)
		if err != nil {
			logError("Failed to get winners", err)
			return
		} else {
			fmt.Print("\n<<<<<<<<<<<<<<<<<<<<<<<< POOLS OF " + player.Name + " >>>>>>>>>>>>>>>>>>>>>>>>>>")
			jsonData, err := json.MarshalIndent(getWinnersResponse, "", "    ")
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Printf("\n%s\n\n", jsonData)
		}

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
