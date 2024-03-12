package goethereumhelper

import (
	"errors"
	// "log"

	"github.com/ethereum/go-ethereum/ethclient"
)

var Client *ethclient.Client
var Clientws *ethclient.Client

/*
GetCustomNetworkClient connects and return a client to user defined Ethereum network
*/
func GetCustomNetworkClient(URL string) (*ethclient.Client, error) {
	if Client != nil {
		return Client, nil
	}
	var err error
	Client, err = ethclient.Dial(URL)
	if err != nil {
		// log.Printf("There was a failure connecting to %s: %+v", URL, err)
		return nil, errors.Join(errors.New("[SubLogs] There was a failure connecting to ethclient"), err)
		// return
	}
	return Client, nil
}

/*
GetCustomNetworkClientWebsocket connects via websocket and return a client to user defined Ethereum network
*/
func GetCustomNetworkClientWebsocket(URL string) (*ethclient.Client, error) {
	if Clientws != nil {
		return Clientws, nil
	}
	var err error
	Clientws, err = ethclient.Dial(URL)
	if err != nil {
		// log.Printf("There was a failure connecting to %s via Websocket: %+v", URL, err)
		return nil, errors.Join(errors.New("[SubLogs] There was a failure connecting to ethclient"), err)

	}
	return Clientws, nil
}

func GetNodeUrl(network int) (string, error) {
	var NodeUrl string
	switch network {
	case 1:
		NodeUrl = "https://ethereum.publicnode.com"
	case 42161:
		NodeUrl = "https://arbitrum-one.publicnode.com"
	case 11155111:
		NodeUrl = "https://ethereum-sepolia.publicnode.com"
	default:
		NodeUrl = "https://ethereum-sepolia.publicnode.com"
	}
	return NodeUrl, nil
}
