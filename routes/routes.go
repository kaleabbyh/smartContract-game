package routes

import (
	"cicada/web-service-gin/config"

	contract "cicada/web-service-gin/controllers/contractController"
	play "cicada/web-service-gin/controllers/gameController"
	playerscore "cicada/web-service-gin/controllers/playerScoreController"
	pool "cicada/web-service-gin/controllers/poolController"

	"cicada/web-service-gin/middlewares"
	"cicada/web-service-gin/utils"

	"os"

	"github.com/gin-gonic/gin"
)

func AllRoutes(router *gin.Engine) {

	db := config.ConnectDB()

	var SecretKey string
	SECRET_KEY_jsonBytes, err := os.ReadFile(os.Getenv("SECRET_KEY"))
	if err != nil {
		// fmt.Printf("Failed to read file: %s", err.Error())
		SecretKey = os.Getenv("SECRET_KEY")
	} else {
		SecretKey = string(SECRET_KEY_jsonBytes)
	}

	var NEXTAUTH_SECRET string
	NEXTAUTH_SECRET_jsonBytes, err := os.ReadFile(os.Getenv("NEXTAUTH_SECRET"))
	if err != nil {
		// fmt.Printf("Failed to read file: %s", err.Error())
		NEXTAUTH_SECRET = os.Getenv("NEXTAUTH_SECRET")
	} else {
		NEXTAUTH_SECRET = string(NEXTAUTH_SECRET_jsonBytes)
	}

	ServerConfigs := utils.ServerConfigs{
		DB:              db,
		SecretKey:       SecretKey,
		NEXTAUTH_SECRET: NEXTAUTH_SECRET,
	}

	poolUserRoute := router.Group("/pool")
	poolUserRoute.Use(middlewares.IsUserAuthenticated(ServerConfigs))
	poolUserRoute.POST("/create-pool", pool.CreatePool(ServerConfigs))
	poolUserRoute.GET("/get-all-pools", pool.GetAllPools(ServerConfigs))
	poolUserRoute.GET("/get-pool/:id", pool.GetPoolByID(ServerConfigs))
	poolUserRoute.POST("/sort-pool", pool.SortPool())
	poolUserRoute.GET("/get-poollis/:poolID/:player", pool.GetAllPoollisPlayerIsIn(ServerConfigs))

	poolAdminRoute := router.Group("/pool")
	poolAdminRoute.Use(middlewares.IsAdminAuthenticated(ServerConfigs))
	poolAdminRoute.PUT("/update-pool/:id", pool.UpdatePool(ServerConfigs))
	poolAdminRoute.DELETE("/delete-pool/:id", pool.DeletePoolWithPoollisByID(ServerConfigs))

	gameRoute := router.Group("/play")
	gameRoute.Use(middlewares.IsUserAuthenticated(ServerConfigs))
	gameRoute.POST("/winners", play.PostWinners())
	gameRoute.GET("/get-winners/:gameid", play.GetWinner(ServerConfigs))

	playerAdminRoute := router.Group("/player")
	playerAdminRoute.Use(middlewares.IsAdminAuthenticated(ServerConfigs))
	playerAdminRoute.POST("/create-player-score", playerscore.CreatePlayer(ServerConfigs))
	playerAdminRoute.PUT("/update-player-score/:id", playerscore.UpdatePlayer(ServerConfigs))
	playerAdminRoute.PUT("/update-player/:gameid", playerscore.UpdateAllPlayerGameid(ServerConfigs))

	playerUserRoute := router.Group("/player")
	playerUserRoute.Use(middlewares.IsUserAuthenticated(ServerConfigs))
	playerUserRoute.GET("/get-player-scores/:gameid", playerscore.GetPlayerScoreByGameID(ServerConfigs))

	ContractRoute := router.Group("/contract")
	ContractRoute.Use(middlewares.IsAdminOrUserAuthenticated(ServerConfigs))
	ContractRoute.GET("/:contract_address", contract.GetContractByContractAddress(ServerConfigs))

	escrowRoute := router.Group("/escrow")
	escrowRoute.Use(middlewares.IsUserAuthenticated(ServerConfigs))
	escrowRoute.POST("/deploy", contract.DeployContract(ServerConfigs))
	escrowRoute.POST("/deposit-eth", contract.DepositETH(ServerConfigs))
	escrowRoute.POST("/deposit-token", contract.DepositToken(ServerConfigs))
	escrowRoute.GET("/isWhitelisted/:network/:contractAddress/:symbol", contract.IsWhitelisted(ServerConfigs))
	// TODO: whitelist will be IsAdminOrUserAuthenticated
	escrowUserARoute := router.Group("/escrow")
	escrowUserARoute.Use(middlewares.IsAdminOrUserAuthenticated(ServerConfigs))
	escrowUserARoute.POST("/whitelist/:chosenTable", contract.WhitelistToken(ServerConfigs))
	escrowUserARoute.GET("/close/:chosenTable/:network/:contractAddress", contract.Close(ServerConfigs))
	escrowUserARoute.GET("/beneficiary-withdraw/:chosenTable/:network/:contractAddress", contract.BeneficiaryWithdraw(ServerConfigs))
	escrowUserARoute.GET("/beneficiary-wd-token/:chosenTable/:network/:contractAddress/:symbol", contract.BeneficiaryWithdrawToken(ServerConfigs))

	// TODO: Only Admin access
	escrowAdminRoute := router.Group("/escrow")
	escrowAdminRoute.Use(middlewares.IsAdminAuthenticated(ServerConfigs))
	escrowAdminRoute.GET("/enableRefunds/:network/:contractAddress", contract.EnableRefunds(ServerConfigs))
	escrowAdminRoute.POST("/withdraw", contract.Withdraw(ServerConfigs))
	escrowAdminRoute.POST("/withdraw-token", contract.WithdrawToken(ServerConfigs))
	//demo contract creare-route
	//ContractRoute.POST("/creat-contract", contract.CreateContract(ServerConfigs))

}
