package goethereumhelper

import (
	"cicada/web-service-gin/config"
	"context"

	// "errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	auth        *bind.TransactOpts
	ks          *keystore.KeyStore
	account     accounts.Account
	fromAddress *common.Address
	err         error
)

type GetGasReq struct {
	ChainId *big.Int
	GasTip *big.Int
	Nonce uint64
	MaxGasFeeAccepted *big.Int
}

func GetAuth(Client *ethclient.Client) (*bind.TransactOpts, error) {
	if auth == nil {
		ks, err = config.ImportKs()
		if err != nil {
			return nil, err
		}
		a1 := accounts.Account{}
		a1.Address = common.HexToAddress(config.Admin)
		account, err = ks.Find(a1)
		if err != nil {
			return nil, err
		}
		fromAddress = &a1.Address

	}

	nonce, err := Client.PendingNonceAt(context.Background(), *fromAddress)
	if err != nil {
		return nil, err
	}
	gasPrice, err := Client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	chainid, err := Client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyStoreTransactorWithChainID(ks, account, chainid)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Estimated gasPrice: %s \n", gasPrice)
	auth.Nonce = big.NewInt(int64(nonce))
	return auth, nil
}

func GetGas(Client *ethclient.Client,sender common.Address) (*GetGasReq,error) {
	chainId, err := Client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	gasTip, err := Client.SuggestGasTipCap(context.Background())
	if err != nil {
		return nil, err
	}

	nonce, err := Client.PendingNonceAt(context.Background(), sender)
	if err != nil {
		return nil, err
	}
	latestEthBlockHeader, err := Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	maxGasFeeAccepted := new(big.Int).Add(
		latestEthBlockHeader.BaseFee,
		gasTip)

	return &GetGasReq{ChainId: chainId, GasTip: gasTip, Nonce: nonce, MaxGasFeeAccepted: maxGasFeeAccepted}, nil
}
