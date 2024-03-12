package controllers

import (
	"cicada/web-service-gin/contracts"
	erc20 "cicada/web-service-gin/contracts/token"
	"cicada/web-service-gin/goethereumhelper"
	"cicada/web-service-gin/models"
	repo "cicada/web-service-gin/repositories"
	docs "cicada/web-service-gin/req-res-models"
	"cicada/web-service-gin/utils"
	"errors"
	"math/big"
	"net/http"
	"strings"

	"context"
	// "fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type TransactionHashResponse struct {
	TransactionHash string
}

type DeployRequest struct {
	Beneficiary  string `binding:"required,eth_addr"`
	InitialOwner string `binding:"required,eth_addr"`
	Network      int
}

type DeployResponse struct {
	ContractAddress string `binding:"required,eth_addr"`
	TransactionHash string
}

type DepositETHRequest struct {
	Contract string `binding:"required,eth_addr"`
	Sender   string `binding:"required,eth_addr"`
	Refundee string `binding:"required,eth_addr"`
	Amount   string
	Network  int
}

type DepositETHResponse struct {
	UnsignedHash string
}

type WhitelistTokenRequest struct {
	Contract     string `binding:"required,eth_addr"`
	TokenAddress string `binding:"required,eth_addr"`
	Symbol       string
	Network      int
}

type WithdrawTokenRequest struct {
	Contract string `binding:"required,eth_addr"`
	Payee    string `binding:"required,eth_addr"`
	Symbol   string
	Network  int
}

type WithdrawRequest struct {
	Contract string `binding:"required,eth_addr"`
	Payee    string `binding:"required,eth_addr"`
	Network  int
}
type DepositTokenRequest struct {
	Contract     string `binding:"required,eth_addr"`
	Sender       string `binding:"required,eth_addr"`
	Refundee     string `binding:"required,eth_addr"`
	TokenAddress string `binding:"required,eth_addr"`
	Amount       string
	Symbol       string
	Network      int
}

type DepositTokenResponse struct {
	UnsignedHash string
	State        string
}

type Params struct {
	ContractAddress string `uri:"contractAddress" binding:"required,eth_addr"`
	Network         int    `uri:"network" binding:"required"`
}

func GetContractByContractAddress(serverConfigs utils.ServerConfigs) func(*gin.Context) {
	return func(c *gin.Context) {

		logger := log.Logger.With().Str("func", "GetContractByContractAddress").Logger()
		var contract *models.EscrowHoF

		id := c.Param("contract_address")
		user, _ := c.Get("user")
		userID, _ := c.Get("userID")

		userID = userID.(string)
		user = user.(string)

		GetContractByContractAddress, err := repo.GetContractByContractAddress(serverConfigs.DB, id)
		if err != nil {
			logger.Error().Err(err).Msgf("Failed to get Contract by Contract-Address: %s", err.Error())

			errorResponse := docs.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Failed to get contract by Contract-Address",
			}

			c.JSON(http.StatusInternalServerError, errorResponse)
			return
		}
		if user == "admin" {
			// contract = GetContractByContractAddress
		} else if userID == GetContractByContractAddress {
			// contract = GetContractByContractAddress
		} else {
			logger.Error().Err(err).Msgf("UnAuthorized")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "UnAuthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":   "success",
			"contract": contract,
		})
	}
}

func CreateContract(serverConfigs utils.ServerConfigs) func(*gin.Context) {
	return func(c *gin.Context) {

		logger := log.Logger.With().Str("func", "CreatePlayer").Logger()
		var requestBody models.EscrowHoF

		err := c.BindJSON(&requestBody)
		if err != nil {
			logger.Error().Err(err).Msg("Invalid request payload")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		contract, err := repo.CreateContract(serverConfigs.DB, &requestBody)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create contract"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"contract": contract,
		})
	}
}

// DeployContract deploys contract.
// @Summary DeployContract
// @Description Used for contract deployment.
// @Tags contracts
// @Accept json
// @Produce json
// @Security bearer
// @Param Authorization header string true "Bearer "
// @Param body body DeployRequest true "Deploy Request"
// @Success 200 {object} DeployResponse
// @Failure 401 {object} docs.ErrorResponse
// @Failure 500 {object} docs.ErrorResponse
// @Router /escrow/deploy [post]
func DeployContract(serverConfigs utils.ServerConfigs) func(*gin.Context) {
	return func(c *gin.Context) {

		logger := log.Logger.With().Str("func", "DeployContract").Logger()
		var requestBody DeployRequest

		err := c.BindJSON(&requestBody)
		if err != nil {
			logger.Error().Err(err).Msg("Invalid request payload")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		network := requestBody.Network
		admin := common.HexToAddress(requestBody.InitialOwner)
		beneficiary := common.HexToAddress(requestBody.Beneficiary)
		initialOwner := common.HexToAddress(requestBody.InitialOwner)

		NodeUrl, _ := goethereumhelper.GetNodeUrl(network)

		Client, err := goethereumhelper.GetCustomNetworkClient(NodeUrl)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to connect client")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to client"})
			return
		}
		result, err := CallDeployContract(Client, beneficiary, initialOwner)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to deploy contract")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create contract"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"contractAddress": result.ContractAddress,
			"transactionHash": result.TransactionHash,
			"adminAddress":    admin,
		})
	}
}

// DepositETH deposits ETH.
// @Summary DepositETH
// @Description Used for depositing ETH.
// @Tags contracts
// @Accept json
// @Produce json
// @Security bearer
// @Param Authorization header string true "Bearer "
// @Param body body DepositETHRequest true "Deposit ETH"
// @Success 200 {object} DepositETHResponse
// @Failure 401 {object} docs.ErrorResponse
// @Failure 500 {object} docs.ErrorResponse
// @Router /escrow/deposit-eth [post]
func DepositETH(serverConfigs utils.ServerConfigs) func(*gin.Context) {
	return func(c *gin.Context) {

		logger := log.Logger.With().Str("func", "DepositETH").Logger()
		var requestBody DepositETHRequest

		err := c.BindJSON(&requestBody)
		if err != nil {
			logger.Error().Err(err).Msg("Invalid request payload")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		network := requestBody.Network
		amount := requestBody.Amount
		contractAddress := common.HexToAddress(requestBody.Contract)
		sender := common.HexToAddress(requestBody.Sender)
		refundee := common.HexToAddress(requestBody.Refundee)

		NodeUrl, _ := goethereumhelper.GetNodeUrl(network)

		Client, err := goethereumhelper.GetCustomNetworkClient(NodeUrl)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to connect client")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to client"})
			return
		}
		result, err := CallDepositETH(Client, contractAddress, amount, sender, refundee)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to run deposit func")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to deposit"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"address":      contractAddress,
			"unsignedHash": result.UnsignedHash,
		})
	}
}

// WhitelistToken Whitelist token.
// @Summary WhitelistToken
// @Description To whitelist a token by symbol.
// @Tags contracts
// @Accept json
// @Produce json
// @Security bearer
// @Param Authorization header string true "Bearer "
// @Param body body WhitelistTokenRequest true "WhitelistToken Request"
// @Success 200 {object} TransactionHashResponse
// @Failure 401 {object} docs.ErrorResponse
// @Failure 500 {object} docs.ErrorResponse
// @Router /escrow/whitelist/{chosenTable} [post]
func WhitelistToken(serverConfigs utils.ServerConfigs) func(*gin.Context) {
	return func(c *gin.Context) {

		logger := log.Logger.With().Str("func", "WhitelistToken").Logger()
		var requestBody WhitelistTokenRequest

		err := c.BindJSON(&requestBody)
		if err != nil {
			logger.Error().Err(err).Msg("Invalid request payload")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		network := requestBody.Network
		symbol := requestBody.Symbol
		contractAddress := common.HexToAddress(requestBody.Contract)
		tokenAddress := common.HexToAddress(requestBody.TokenAddress)

		id := requestBody.Contract
		chosenTable := c.Param("chosenTable")
		user, _ := c.Get("user")
		userID, _ := c.Get("userID")

		userID = userID.(string)
		user = user.(string)
		if user != "admin" {

			GetUserIdByContractAddress, err := repo.GetUserIdByContractAddress(serverConfigs.DB, id, chosenTable)
			if err != nil {
				logger.Error().Err(err).Msgf("Failed to get Contract-Address from DB: %s", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to authorize, DB error"})
				return
			}
			var UserIdComp string
			switch chosenTable {
			case "hof":
				UserIdComp = GetUserIdByContractAddress.ContractHoF.UserId
			case "hire":
				UserIdComp = GetUserIdByContractAddress.ContractHire.UserId
			default:
				logger.Error().Err(err).Msgf("Failed to find escrow table")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find escrow table"})
				return
			}

			if userID != UserIdComp {
				logger.Error().Err(err).Msgf("UnAuthorized access")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "UnAuthorized access"})
				return
			}

		}

		NodeUrl, _ := goethereumhelper.GetNodeUrl(network)

		Client, err := goethereumhelper.GetCustomNetworkClient(NodeUrl)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to connect client")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to client"})
			return
		}
		whitelisted, err := CallIsWhitelisted(Client, contractAddress, symbol, "check")
		if err != nil {
			logger.Error().Err(err).Msg("Failed to check if Whitelisted")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check if Whitelisted"})
			return
		}
		zero := common.Address{}
		if zero != *whitelisted {
			logger.Error().Err(err).Msg("Already whitelisted address")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Token already whitelisted"})
			return
		}
		result, err := CallWhitelistToken(Client, contractAddress, symbol, tokenAddress)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to whitelist the token")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to whitelist"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"transactionHash": result.Hash().Hex(),
		})
	}
}

// DepositToken deposits token.
// @Summary DepositToken
// @Description Used for depositing token.
// @Tags contracts
// @Accept json
// @Produce json
// @Security bearer
// @Param Authorization header string true "Bearer "
// @Param body body DepositTokenRequest true "Deposit token"
// @Success 200 {object} DepositTokenResponse
// @Failure 401 {object} docs.ErrorResponse
// @Failure 500 {object} docs.ErrorResponse
// @Router /escrow/deposit-token [post]
func DepositToken(serverConfigs utils.ServerConfigs) func(*gin.Context) {
	return func(c *gin.Context) {

		logger := log.Logger.With().Str("func", "DepositToken").Logger()
		var requestBody DepositTokenRequest

		err := c.BindJSON(&requestBody)
		if err != nil {
			logger.Error().Err(err).Msg("Invalid request payload")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		network := requestBody.Network
		amount := requestBody.Amount
		symbol := requestBody.Symbol
		contractAddress := common.HexToAddress(requestBody.Contract)
		sender := common.HexToAddress(requestBody.Sender)
		tokenAddress := common.HexToAddress(requestBody.TokenAddress)
		refundee := common.HexToAddress(requestBody.Refundee)

		NodeUrl, _ := goethereumhelper.GetNodeUrl(network)

		Client, err := goethereumhelper.GetCustomNetworkClient(NodeUrl)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to connect client")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to client"})
			return
		}
		result, err := CallDepositToken(Client, contractAddress, amount, symbol, sender, refundee, tokenAddress)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to run depositToken func")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to deposit token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"address":      contractAddress,
			"unsignedHash": result.UnsignedHash,
			"state":        result.State,
		})
	}
}

// Close change state to close.
// @Summary Close
// @Description Changes state of contract to closed.
// @Tags contracts
// @Accept json
// @Produce json
// @Param chosenTable path string true "Chosen Table"
// @Param network path string true "network"
// @Param contractAddress path string true "contract address"
// @Security bearer
// @Param Authorization header string true "Bearer "
// @Success 200 {object} TransactionHashResponse
// @Failure 401 {object} docs.ErrorResponse
// @Failure 500 {object} docs.ErrorResponse
// @Router /escrow/close/{chosenTable}/{network}/{contractAddress} [get]
func Close(serverConfigs utils.ServerConfigs) func(*gin.Context) {
	return func(c *gin.Context) {

		logger := log.Logger.With().Str("func", "Close").Logger()
		var params Params

		err := c.ShouldBindUri(&params)
		if err != nil {
			logger.Error().Err(err).Msg("Invalid request parameter")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request param"})
			return
		}

		network := params.Network
		contractAddress := common.HexToAddress(params.ContractAddress)

		id := params.ContractAddress
		chosenTable := c.Param("chosenTable")
		user, _ := c.Get("user")
		userID, _ := c.Get("userID")

		userID = userID.(string)
		user = user.(string)
		if user != "admin" {

			GetUserIdByContractAddress, err := repo.GetUserIdByContractAddress(serverConfigs.DB, id, chosenTable)
			if err != nil {
				logger.Error().Err(err).Msgf("Failed to get Contract-Address from DB: %s", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to authorize, DB error"})
				return
			}
			var UserIdComp string
			switch chosenTable {
			case "hof":
				UserIdComp = GetUserIdByContractAddress.ContractHoF.UserId
			case "hire":
				UserIdComp = GetUserIdByContractAddress.ContractHire.UserId
			default:
				logger.Error().Err(err).Msgf("Failed to find escrow table")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find escrow table"})
				return
			}

			if userID != UserIdComp {
				logger.Error().Err(err).Msgf("UnAuthorized access")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "UnAuthorized access"})
				return
			}
		}

		NodeUrl, _ := goethereumhelper.GetNodeUrl(network)

		Client, err := goethereumhelper.GetCustomNetworkClient(NodeUrl)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to connect client")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to client"})
			return
		}
		result, err := CallClose(Client, contractAddress)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to change state to close")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to close"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"transactionHash": result.Hash().Hex(),
		})
	}
}

// EnableRefunds change state to Refund.
// @Summary EnableRefunds
// @Description Changes state of contract to refund.
// @Tags contracts
// @Accept json
// @Produce json
// @Param network path string true "network"
// @Param contractAddress path string true "contract address"
// @Security bearer
// @Param Authorization header string true "Bearer "
// @Success 200 {object} TransactionHashResponse
// @Failure 401 {object} docs.ErrorResponse
// @Failure 500 {object} docs.ErrorResponse
// @Router /escrow/enableRefunds/{network}/{contractAddress} [get]
func EnableRefunds(serverConfigs utils.ServerConfigs) func(*gin.Context) {
	return func(c *gin.Context) {

		logger := log.Logger.With().Str("func", "EnableRefunds").Logger()
		var params Params

		err := c.ShouldBindUri(&params)
		if err != nil {
			logger.Error().Err(err).Msg("Invalid request parameter")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request param"})
			return
		}

		network := params.Network
		contractAddress := common.HexToAddress(params.ContractAddress)

		NodeUrl, _ := goethereumhelper.GetNodeUrl(network)

		Client, err := goethereumhelper.GetCustomNetworkClient(NodeUrl)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to connect client")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to client"})
			return
		}
		result, err := CallEnableRefunds(Client, contractAddress)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to change state to refund")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enable refund"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"transactionHash": result.Hash().Hex(),
		})
	}
}

// BeneficiaryWithdraw transfer to beneficiary.
// @Summary BeneficiaryWithdraw
// @Description Transfer asset to beneficiary.
// @Tags contracts
// @Accept json
// @Produce json
// @Param chosenTable path string true "Chosen Table"
// @Param network path string true "network"
// @Param contractAddress path string true "contract address"
// @Security bearer
// @Param Authorization header string true "Bearer "
// @Success 200 {object} TransactionHashResponse
// @Failure 401 {object} docs.ErrorResponse
// @Failure 500 {object} docs.ErrorResponse
// @Router /escrow/beneficiary-withdraw/{chosenTable}/{network}/{contractAddress} [get]
func BeneficiaryWithdraw(serverConfigs utils.ServerConfigs) func(*gin.Context) {
	return func(c *gin.Context) {

		logger := log.Logger.With().Str("func", "BeneficiaryWithdraw").Logger()
		var params Params

		err := c.ShouldBindUri(&params)
		if err != nil {
			logger.Error().Err(err).Msg("Invalid request parameter")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request param"})
			return
		}

		network := params.Network
		contractAddress := common.HexToAddress(params.ContractAddress)

		id := params.ContractAddress
		chosenTable := c.Param("chosenTable")
		user, _ := c.Get("user")
		userID, _ := c.Get("userID")

		userID = userID.(string)
		user = user.(string)
		if user != "admin" {

			GetUserIdByContractAddress, err := repo.GetUserIdByContractAddress(serverConfigs.DB, id, chosenTable)
			if err != nil {
				logger.Error().Err(err).Msgf("Failed to get Contract-Address from DB: %s", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to authorize, DB error"})
				return
			}
			var UserIdComp string
			switch chosenTable {
			case "hof":
				UserIdComp = GetUserIdByContractAddress.ContractHoF.UserId
			case "hire":
				UserIdComp = GetUserIdByContractAddress.ContractHire.UserId
			default:
				logger.Error().Err(err).Msgf("Failed to find escrow table")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find escrow table"})
				return
			}

			if userID != UserIdComp {
				logger.Error().Err(err).Msgf("UnAuthorized access")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "UnAuthorized access"})
				return
			}
		}

		NodeUrl, _ := goethereumhelper.GetNodeUrl(network)

		Client, err := goethereumhelper.GetCustomNetworkClient(NodeUrl)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to connect client")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to client"})
			return
		}
		balance, err := Client.BalanceAt(context.Background(), contractAddress, nil)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to fetch balance of contract")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch balance of contract"})
			return
		}
		if balance.Cmp(big.NewInt(0)) == 0 {
			logger.Error().Err(err).Msg("restricting benef-withdraw, eth balance is zero")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Not allowed, eth balance is zero"})
			return
		}
		result, err := CallBeneficiaryWithdraw(Client, contractAddress)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to Withdraw for beneficiary")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Withdraw for beneficiary"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"transactionHash": result.Hash().Hex(),
		})
	}
}

// Withdraw transfer to payee.
// @Summary Withdraw
// @Description Transfer asset to payee.
// @Tags contracts
// @Accept json
// @Produce json
// @Security bearer
// @Param Authorization header string true "Bearer "
// @Param body body WithdrawRequest true "Withdraw eth"
// @Success 200 {object} TransactionHashResponse
// @Failure 401 {object} docs.ErrorResponse
// @Failure 500 {object} docs.ErrorResponse
// @Router /escrow/withdraw [post]
func Withdraw(serverConfigs utils.ServerConfigs) func(*gin.Context) {
	return func(c *gin.Context) {

		logger := log.Logger.With().Str("func", "Withdraw").Logger()
		var requestBody WithdrawRequest

		err := c.BindJSON(&requestBody)
		if err != nil {
			logger.Error().Err(err).Msg("Invalid request payload")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		network := requestBody.Network
		contractAddress := common.HexToAddress(requestBody.Contract)
		payee := common.HexToAddress(requestBody.Payee)

		NodeUrl, _ := goethereumhelper.GetNodeUrl(network)

		Client, err := goethereumhelper.GetCustomNetworkClient(NodeUrl)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to connect client")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to client"})
			return
		}
		balance, err := CallDepositsOf(Client, contractAddress, payee)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to get eth-deposit for payee")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get eth-deposit for payee"})
			return
		}
		if balance.Cmp(big.NewInt(0)) == 0 {
			logger.Error().Err(err).Msg("restricting withdraw, eth depositOf payee is zero")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Not allowed, eth deposit of payee is zero"})
			return
		}
		result, err := CallWithdraw(Client, contractAddress, payee)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to Withdraw ETH to payee")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Withdraw ETH to payee"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"transactionHash": result.Hash().Hex(),
		})
	}
}

// BeneficiaryWithdraw-Token transfer to beneficiary.
// @Summary BeneficiaryWithdrawToken
// @Description Transfer specified token symbol to beneficiary.
// @Tags contracts
// @Accept json
// @Produce json
// @Param chosenTable path string true "Chosen Table"
// @Param network path string true "network"
// @Param contractAddress path string true "contract address"
// @Param symbol path string true "symbol"
// @Security bearer
// @Param Authorization header string true "Bearer "
// @Success 200 {object} TransactionHashResponse
// @Failure 401 {object} docs.ErrorResponse
// @Failure 500 {object} docs.ErrorResponse
// @Router /escrow/beneficiary-wd-token/{chosenTable}/{network}/{contractAddress}/{symbol} [get]
func BeneficiaryWithdrawToken(serverConfigs utils.ServerConfigs) func(*gin.Context) {
	return func(c *gin.Context) {

		logger := log.Logger.With().Str("func", "BeneficiaryWithdrawToken").Logger()
		var params Params

		err := c.ShouldBindUri(&params)
		if err != nil {
			logger.Error().Err(err).Msg("Invalid request parameter")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request param"})
			return
		}

		network := params.Network
		contractAddress := common.HexToAddress(params.ContractAddress)
		symbol := c.Param("symbol")

		id := params.ContractAddress
		chosenTable := c.Param("chosenTable")
		user, _ := c.Get("user")
		userID, _ := c.Get("userID")

		userID = userID.(string)
		user = user.(string)
		if user != "admin" {

			GetUserIdByContractAddress, err := repo.GetUserIdByContractAddress(serverConfigs.DB, id, chosenTable)
			if err != nil {
				logger.Error().Err(err).Msgf("Failed to get Contract-Address from DB: %s", err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to authorize, DB error"})
				return
			}
			var UserIdComp string
			switch chosenTable {
			case "hof":
				UserIdComp = GetUserIdByContractAddress.ContractHoF.UserId
			case "hire":
				UserIdComp = GetUserIdByContractAddress.ContractHire.UserId
			default:
				logger.Error().Err(err).Msgf("Failed to find escrow table")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find escrow table"})
				return
			}

			if userID != UserIdComp {
				logger.Error().Err(err).Msgf("UnAuthorized access")
				c.JSON(http.StatusInternalServerError, gin.H{"error": "UnAuthorized access"})
				return
			}
		}

		NodeUrl, _ := goethereumhelper.GetNodeUrl(network)

		Client, err := goethereumhelper.GetCustomNetworkClient(NodeUrl)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to connect client")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to client"})
			return
		}

		result, err := CallBeneficiaryWithdrawToken(Client, contractAddress, symbol)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to Withdraw-token for beneficiary")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Withdraw-token for beneficiary"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"transactionHash": result.Hash().Hex(),
		})
	}
}

// WithdrawToken transfer to payee.
// @Summary WithdrawToken
// @Description Transfer asset to payee.
// @Tags contracts
// @Accept json
// @Produce json
// @Security bearer
// @Param Authorization header string true "Bearer "
// @Param body body WithdrawTokenRequest true "Withdraw token"
// @Success 200 {object} TransactionHashResponse
// @Failure 401 {object} docs.ErrorResponse
// @Failure 500 {object} docs.ErrorResponse
// @Router /escrow/withdraw-token [post]
func WithdrawToken(serverConfigs utils.ServerConfigs) func(*gin.Context) {
	return func(c *gin.Context) {

		logger := log.Logger.With().Str("func", "WithdrawToken").Logger()
		var requestBody WithdrawTokenRequest

		err := c.BindJSON(&requestBody)
		if err != nil {
			logger.Error().Err(err).Msg("Invalid request payload")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		network := requestBody.Network
		symbol := requestBody.Symbol
		contractAddress := common.HexToAddress(requestBody.Contract)
		payee := common.HexToAddress(requestBody.Payee)

		NodeUrl, _ := goethereumhelper.GetNodeUrl(network)

		Client, err := goethereumhelper.GetCustomNetworkClient(NodeUrl)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to connect client")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to client"})
			return
		}
		balance, err := CallTokenDepositsOf(Client, contractAddress, payee, symbol)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to get token deposit for payee")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get token deposit for payee"})
			return
		}
		if balance.Cmp(big.NewInt(0)) == 0 {
			logger.Error().Err(err).Msg("restricting withdraw, token balance is zero")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Not allowed, Token balance of payee is zero"})
			return
		}
		result, err := CallWithdrawToken(Client, contractAddress, symbol, payee)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to Withdraw the token")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Withdraw token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"transactionHash": result.Hash().Hex(),
		})
	}
}

// IsWhitelisted check if token is whitelisted.
// @Summary BeneficiaryWithdraw
// @Description Checks if token symbol is whitelisted.
// @Tags contracts
// @Accept json
// @Produce json
// @Param network path string true "network"
// @Param contractAddress path string true "contract address"
// @Param symbol path string true "symbol"
// @Security bearer
// @Param Authorization header string true "Bearer "
// @Success 200 {object} TransactionHashResponse
// @Failure 401 {object} docs.ErrorResponse
// @Failure 500 {object} docs.ErrorResponse
// @Router /escrow/isWhitelisted/{network}/{contractAddress}/{symbol} [get]
func IsWhitelisted(serverConfigs utils.ServerConfigs) func(*gin.Context) {
	return func(c *gin.Context) {

		logger := log.Logger.With().Str("func", "IsWhitelisted").Logger()
		var params Params

		err := c.ShouldBindUri(&params)
		if err != nil {
			logger.Error().Err(err).Msg("Invalid request parameter")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request param"})
			return
		}

		network := params.Network
		contractAddress := common.HexToAddress(params.ContractAddress)
		symbol := c.Param("symbol")

		NodeUrl, _ := goethereumhelper.GetNodeUrl(network)

		Client, err := goethereumhelper.GetCustomNetworkClient(NodeUrl)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to connect client")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to client"})
			return
		}
		result, err := CallIsWhitelisted(Client, contractAddress, symbol, "wl")
		if err != nil {
			logger.Error().Err(err).Msg("Failed to run WhitelistedToken function")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to run WhitelistedToken function"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"address": result,
		})
	}
}

func CallDeployContract(Client *ethclient.Client, beneficiary common.Address, initialOwner common.Address) (*DeployResponse, error) {

	auth, err := goethereumhelper.GetAuth(Client)
	if err != nil {
		return nil, err
	}

	address, tx, instance, err := contracts.DeployRefundEscrow(auth, Client, beneficiary, initialOwner)
	if err != nil {
		return nil, err
	}

	_ = address
	_ = instance

	// wait until the contract is deployed
	address, err = bind.WaitDeployed(context.Background(), Client, tx)
	if err != nil {
		return nil, err
	}

	// fmt.Println(address.Hex()) // 0x037D12a94dd5eE285F5efE46820F1e5B2c843355

	return &DeployResponse{ContractAddress: address.Hex(), TransactionHash: tx.Hash().Hex()}, nil
}

func CallDepositETH(Client *ethclient.Client, contractAddress common.Address, amount string, sender common.Address, refundee common.Address) (*DepositETHResponse, error) {

	abiJson := contracts.RefundEscrowABI
	abiP, err := abi.JSON(strings.NewReader(abiJson))
	if err != nil {
		return nil, err
	}

	// publicKeyECDSA, ok := privateKey.Public().(*ecdsa.PublicKey)
	// if !ok {
	// 	return nil, errors.New("PublicKey error")
	// }
	// gasPrice, err := Client.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	return nil, err
	// }

	bytesData, err := abiP.Pack("deposit", refundee)
	if err != nil {
		return nil, err
	}

	i := new(big.Int)
	valueToSend, ok := i.SetString(amount, 10)
	if !ok {
		return nil, err
	}

	val, err := goethereumhelper.GetGas(Client, sender)
	if err != nil {
		return nil, err
	}

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   val.ChainId,
		Nonce:     val.Nonce,
		GasTipCap: val.GasTip,
		GasFeeCap: val.MaxGasFeeAccepted,
		Gas:       310000,
		To:        &contractAddress,
		Data:      bytesData,
		Value:     valueToSend,
	})
	// signedTx, err := types.SignTx(tx, types.NewLondonSigner(tx.ChainId()), privateKey)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Println(signedTx.Hash().Hex())
	// fmt.Println(tx.Hash().Hex())
	// err = Client.SendTransaction(context.TODO(), signedTx)
	// if err != nil {
	// 	return nil, err
	// }
	return &DepositETHResponse{UnsignedHash: tx.Hash().String()}, nil
}

func CallWhitelistToken(Client *ethclient.Client, contractAddress common.Address, symbol string, tokenAddress common.Address) (*types.Transaction, error) {
	// Convert symbol string to byte32
	symbol32 := [32]byte{}
	copy(symbol32[:], []byte(symbol))

	auth, err := goethereumhelper.GetAuth(Client)
	if err != nil {
		return nil, err
	}

	instance, err := contracts.NewRefundEscrow(contractAddress, Client)
	if err != nil {
		return nil, err
	}
	// fmt.Println("Contract loaded")

	tx, err := instance.WhitelistToken(auth, symbol32, tokenAddress)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("tx sent: %s \n", tx.Hash().Hex())
	return tx, nil
}

func CallApproveToken(Client *ethclient.Client, contractAddress common.Address, amount string, sender common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	val, err := goethereumhelper.GetGas(Client, sender)
	if err != nil {
		return nil, err
	}

	i := new(big.Int)
	valueToSend, ok := i.SetString(amount, 10)
	if !ok {
		return nil, err
	}

	// approveFnSignature := []byte("approve(address spender, uint256 amount)")
	// hash := sha3.NewLegacyKeccak256()
	// hash.Write(approveFnSignature)
	// methodID := hash.Sum(nil)[:4]
	// fmt.Println(hexutil.Encode(methodID))

	// paddedAmount := common.LeftPadBytes(valueToSend.Bytes(), 32)
	// fmt.Println(hexutil.Encode(paddedAmount))

	abiJson := "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	abiP, err := abi.JSON(strings.NewReader(abiJson))
	if err != nil {
		return nil, err
	}

	bytesData, err := abiP.Pack("approve", contractAddress, valueToSend)
	if err != nil {
		return nil, err
	}

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   val.ChainId,
		Nonce:     val.Nonce,
		GasTipCap: val.GasTip,
		GasFeeCap: val.MaxGasFeeAccepted,
		Gas:       310000,
		To:        &tokenAddress,
		Data:      bytesData,
	})

	return tx, nil
}

func CallDepositToken(Client *ethclient.Client, contractAddress common.Address, amount string, symbol string, sender common.Address, refundee common.Address, tokenAddress common.Address) (*DepositTokenResponse, error) {
	// Check if token has allowance
	allowance, err := CallAllowance(Client, contractAddress, sender, tokenAddress)
	if err != nil {
		return nil, err
	}

	// Convert string amount to Big.Int
	i := new(big.Int)
	valueToSend, ok := i.SetString(amount, 10)
	if !ok {
		return nil, errors.New("error converting string to bigInt")
	}

	// Check if amount is equal or greater than allowance, ApproveToken if not
	if allowance.Cmp(valueToSend) == -1 {
		tx, err := CallApproveToken(Client, contractAddress, amount, sender, tokenAddress)
		if err != nil {
			return nil, err
		}

		return &DepositTokenResponse{UnsignedHash: tx.Hash().Hex(), State: "approve"}, nil
	}

	abiJson := contracts.RefundEscrowABI
	abiP, err := abi.JSON(strings.NewReader(abiJson))
	if err != nil {
		return nil, err
	}

	// Convert symbol string to byte32
	symbol32 := [32]byte{}
	copy(symbol32[:], []byte(symbol))

	bytesData, err := abiP.Pack("depositToken", valueToSend, symbol32, refundee)
	if err != nil {
		return nil, err
	}

	val, err := goethereumhelper.GetGas(Client, sender)
	if err != nil {
		return nil, err
	}

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   val.ChainId,
		Nonce:     val.Nonce,
		GasTipCap: val.GasTip,
		GasFeeCap: val.MaxGasFeeAccepted,
		Gas:       310000,
		To:        &contractAddress,
		Data:      bytesData,
	})

	return &DepositTokenResponse{UnsignedHash: tx.Hash().Hex(), State: "deposit"}, nil
}

func CallAllowance(Client *ethclient.Client, contractAddress common.Address, owner common.Address, tokenAddress common.Address) (*big.Int, error) {
	instance, err := erc20.NewContractsCaller(tokenAddress, Client)
	if err != nil {
		return nil, err
	}
	// fmt.Println("Contract loaded")

	allowance, err := instance.Allowance(&bind.CallOpts{}, owner, contractAddress)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("Allowance amount: %s", allowance)

	return allowance, nil
}

func CallClose(Client *ethclient.Client, contractAddress common.Address) (*types.Transaction, error) {
	auth, err := goethereumhelper.GetAuth(Client)
	if err != nil {
		return nil, err
	}
	// fmt.Println(contractAddress)
	instance, err := contracts.NewRefundEscrowTransactor(contractAddress, Client)
	if err != nil {
		return nil, err
	}
	// fmt.Println("Contract loaded")

	tx, err := instance.Close(auth)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func CallEnableRefunds(Client *ethclient.Client, contractAddress common.Address) (*types.Transaction, error) {
	auth, err := goethereumhelper.GetAuth(Client)
	if err != nil {
		return nil, err
	}
	// fmt.Println(contractAddress)
	instance, err := contracts.NewRefundEscrowTransactor(contractAddress, Client)
	if err != nil {
		return nil, err
	}
	// fmt.Println("Contract loaded")

	tx, err := instance.EnableRefunds(auth)
	if err != nil {
		return nil, err
	}

	// state, err := instance.State(&bind.CallOpts{})
	// if err != nil {
	// 	return nil, err
	// }

	return tx, nil
}

func CallWithdraw(Client *ethclient.Client, contractAddress common.Address, payee common.Address) (*types.Transaction, error) {
	auth, err := goethereumhelper.GetAuth(Client)
	if err != nil {
		return nil, err
	}

	instance, err := contracts.NewRefundEscrowTransactor(contractAddress, Client)
	if err != nil {
		return nil, err
	}

	tx, err := instance.Withdraw(auth, payee)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func CallWithdrawToken(Client *ethclient.Client, contractAddress common.Address, symbol string, payee common.Address) (*types.Transaction, error) {
	// Convert symbol string to byte32
	symbol32 := [32]byte{}
	copy(symbol32[:], []byte(symbol))

	auth, err := goethereumhelper.GetAuth(Client)
	if err != nil {
		return nil, err
	}

	instance, err := contracts.NewRefundEscrowTransactor(contractAddress, Client)
	if err != nil {
		return nil, err
	}

	tx, err := instance.WithdrawToken(auth, symbol32, payee)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func CallBeneficiaryWithdraw(Client *ethclient.Client, contractAddress common.Address) (*types.Transaction, error) {
	auth, err := goethereumhelper.GetAuth(Client)
	if err != nil {
		return nil, err
	}

	instance, err := contracts.NewRefundEscrowTransactor(contractAddress, Client)
	if err != nil {
		return nil, err
	}

	tx, err := instance.BeneficiaryWithdraw(auth)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func CallBeneficiaryWithdrawToken(Client *ethclient.Client, contractAddress common.Address, symbol string) (*types.Transaction, error) {
	// Convert symbol string to byte32
	symbol32 := [32]byte{}
	copy(symbol32[:], []byte(symbol))

	auth, err := goethereumhelper.GetAuth(Client)
	if err != nil {
		return nil, err
	}

	instance, err := contracts.NewRefundEscrowTransactor(contractAddress, Client)
	if err != nil {
		return nil, err
	}

	tx, err := instance.BeneficiaryTokenWithdraw(auth, symbol32)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func CallIsWhitelisted(Client *ethclient.Client, contractAddress common.Address, symbol string, who string) (*common.Address, error) {
	// Convert symbol string to byte32
	symbol32 := [32]byte{}
	copy(symbol32[:], []byte(symbol))

	instance, err := contracts.NewRefundEscrowCaller(contractAddress, Client)
	if err != nil {
		return nil, err
	}
	if who == "check" {
		state, err := instance.State(&bind.CallOpts{})
		if err != nil {
			return nil, err
		}
		if state != 0 {
			return nil, errors.New("contract is not on active state")
		}
	}

	address, err := instance.WhitelistedToken(&bind.CallOpts{}, symbol32)
	if err != nil {
		return nil, err
	}

	return &address, nil
}

func CallTokenDepositsOf(Client *ethclient.Client, contractAddress common.Address, payee common.Address, symbol string) (*big.Int, error) {
	// Convert symbol string to byte32
	symbol32 := [32]byte{}
	copy(symbol32[:], []byte(symbol))

	instance, err := contracts.NewRefundEscrowCaller(contractAddress, Client)
	if err != nil {
		return nil, err
	}

	balance, err := instance.TokenDepositsOf(&bind.CallOpts{}, payee, symbol32)
	if err != nil {
		return nil, err
	}

	return balance, nil
}

func CallDepositsOf(Client *ethclient.Client, contractAddress common.Address, payee common.Address) (*big.Int, error) {
	instance, err := contracts.NewRefundEscrowCaller(contractAddress, Client)
	if err != nil {
		return nil, err
	}

	balance, err := instance.DepositsOf(&bind.CallOpts{}, payee)
	if err != nil {
		return nil, err
	}

	return balance, nil
}
