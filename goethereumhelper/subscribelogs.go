package goethereumhelper

import (
	"context"
	"errors"
	"log"
	"sync"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// SubLogs Subscribe to watch to notifications to a specific Ethereum address
func SubLogs(addressToWatch common.Address, wg *sync.WaitGroup) error {
	log.Println("[SubLogs] Waiting network confirmation to address ", addressToWatch.String(), " ...")
	defer wg.Done()
	wsClient, err := GetCustomNetworkClientWebsocket("<<put in here your EVM Node URL>>")
	if err != nil {
		return errors.Join(errors.New("[SubLogs] Failed to connect via WS"), err)
	}
	query := ethereum.FilterQuery{
		Addresses: []common.Address{addressToWatch},
	}
	logs := make(chan types.Log)
	sub, err := wsClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		return errors.Join(errors.New("[SubLogs] Failed to subscribe to events on the network"), err)
	}

	for {
		select {
		case err := <-sub.Err():
			return errors.Join(errors.New("[SubLogs] Error from network connection"), err)
		case infoLog := <-logs:
			log.Println("[SubLogs] Information received from Rinkeby. Address: ", addressToWatch.String(), " - Information: ", infoLog)
			return nil
		}
	}
}
