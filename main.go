package main

import (
	"cicada/web-service-gin/config"
	"cicada/web-service-gin/routes"

	docs "cicada/web-service-gin/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	config.LoadConfig()
}

// @title cicada REST API
// @version 2.0
// @description cicada REST API description
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	gin.SetMode(gin.ReleaseMode)

	logger := log.Logger.With().Str("func", "main").Logger()
	logger.Info().Msg("Server starting...")

	server := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://cicada-frontend.com"}
	// config.AllowAllOrigins = true
	server.Use(cors.New(config))

	docs.SwaggerInfo.Schemes = []string{"http"}
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.AllRoutes(server)
	err := server.Run(":8080")

	if err != nil {
		logger.Error().Err(err).Msg("Server failed to start")
	}
}
