package middlewares

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"cicada/web-service-gin/repositories"
	"cicada/web-service-gin/utils"

	"github.com/gin-gonic/gin"

	"github.com/rs/zerolog/log"
)

// type UserPayload struct {
// 	Name          string `json:"name"`
// 	Email         string `json:"email"`
// 	Picture       string `json:"picture"`
// 	Sub           string `json:"sub"`
// 	IsAuditor     bool   `json:"isAuditor"`
// 	Image         string `json:"image"`
// 	WalletAddress string `json:"wallet_address"`
// 	ChainID       string `json:"chainId"`
// 	Username      string `json:"username"`
// 	AccessToken   string `json:"access_token"`
// 	Iat           int64  `json:"iat"`
// 	Exp           int64  `json:"exp"`
// 	Jti           string `json:"jti"`
// }

//	type AdminPayload struct {
//		Email         string `json:"email"`
//		Picture       string `json:"picture"`
//		Sub           string `json:"sub"`
//		Image         string `json:"image"`
//		WalletAddress string `json:"wallet_address"`
//		ChainID       int    `json:"chainId"`
//		Username      string `json:"username"`
//		Iat           int64  `json:"iat"`
//		Exp           int64  `json:"exp"`
//		Jti           string `json:"jti"`
//	}
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func IsUserAuthenticated(serverConfigs utils.ServerConfigs) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info().Msg("auth middleware")
		authHeader := c.GetHeader("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")

		encryptionSecret, err := utils.GetDerivedEncryptionKey(serverConfigs.NEXTAUTH_SECRET, "")
		if err != nil {
			log.Error().Err(err).Msg("not authorized")
			errorResponse := ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "not authorized",
			}
			c.JSON(http.StatusUnauthorized, errorResponse)
			c.Abort()
			return
		}

		decryptedPayload, err := utils.DecodeAndDecryptToken(token, encryptionSecret)
		if err != nil {
			log.Error().Err(err).Msg("not authorized")
			errorResponse := ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "not authorized",
			}
			c.JSON(http.StatusUnauthorized, errorResponse)
			c.Abort()
			return
		}

		//var payload UserPayload
		// err = json.Unmarshal([]byte(decryptedPayload), &payload)
		// if err != nil {
		// 	log.Error().Err(err).Msg("not authorized")
		// 	errorResponse := ErrorResponse{
		// 		Code:    http.StatusUnauthorized,
		// 		Message: "not authorized",
		// 	}
		// 	c.JSON(http.StatusUnauthorized, errorResponse)
		// 	c.Abort()
		// 	return
		// }

		// currentTime := time.Now().Unix()
		// if payload.Exp <= currentTime {
		// 	log.Error().Err(err).Msg("Expired token")
		// 	errorResponse := ErrorResponse{
		// 		Code:    http.StatusUnauthorized,
		// 		Message: "Expired token",
		// 	}
		// 	c.JSON(http.StatusUnauthorized, errorResponse)
		// 	c.Abort()
		// 	return
		// }

		var payload map[string]interface{}
		err = json.Unmarshal([]byte(decryptedPayload), &payload)
		if err != nil {
			log.Error().Err(err).Msg("not authorized")
			errorResponse := ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "not authorized",
			}
			c.JSON(http.StatusUnauthorized, errorResponse)
			c.Abort()
			return
		}

		currentTime := time.Now().Unix()
		exp, ok := payload["exp"].(float64)
		if !ok {
			log.Error().Err(err).Msg("Invalid token expiration")
			errorResponse := ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "Invalid token expiration",
			}
			c.JSON(http.StatusUnauthorized, errorResponse)
			c.Abort()
			return
		}

		if int64(exp) <= currentTime {
			log.Error().Err(err).Msg("Expired token")
			errorResponse := ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "Expired token",
			}
			c.JSON(http.StatusUnauthorized, errorResponse)
			c.Abort()
			return
		}

		c.Next()
	}
}

func IsAdminAuthenticated(serverConfigs utils.ServerConfigs) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info().Msg("admin auth middleware")

		authHeader := c.GetHeader("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")

		encryptionSecret, err := utils.GetDerivedEncryptionKey(serverConfigs.NEXTAUTH_SECRET, "")
		if err != nil {
			log.Error().Err(err).Msg("not authorized")
			errorResponse := ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "not authorized",
			}
			c.JSON(http.StatusUnauthorized, errorResponse)
			c.Abort()
			return
		}

		decryptedPayload, err := utils.DecodeAndDecryptToken(token, encryptionSecret)
		if err != nil {
			log.Error().Err(err).Msg("not authorized")
			errorResponse := ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "not authorized",
			}
			c.JSON(http.StatusUnauthorized, errorResponse)
			c.Abort()
			return
		}
		log.Info().Msgf("payload successfully decrypted")

		// var payload AdminPayload
		// err = json.Unmarshal([]byte(decryptedPayload), &payload)
		// if err != nil {
		// 	log.Error().Err(err).Msg("not authorized")
		// 	errorResponse := ErrorResponse{
		// 		Code:    http.StatusUnauthorized,
		// 		Message: "not authorized",
		// 	}
		// 	c.JSON(http.StatusUnauthorized, errorResponse)
		// 	c.Abort()
		// 	return
		// }

		// currentTime := time.Now().Unix()
		// if payload.Exp <= currentTime {
		// 	log.Error().Err(err).Msg("Expired token")
		// 	errorResponse := ErrorResponse{
		// 		Code:    http.StatusUnauthorized,
		// 		Message: "Expired token",
		// 	}
		// 	c.JSON(http.StatusUnauthorized, errorResponse)
		// 	c.Abort()
		// 	return
		// }

		var payload map[string]interface{}
		err = json.Unmarshal([]byte(decryptedPayload), &payload)
		if err != nil {
			log.Error().Err(err).Msg("not authorized")
			errorResponse := ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "not authorized",
			}
			c.JSON(http.StatusUnauthorized, errorResponse)
			c.Abort()
			return
		}

		currentTime := time.Now().Unix()
		exp, ok := payload["exp"].(float64)
		if !ok {
			log.Error().Err(err).Msg("Invalid token expiration")
			errorResponse := ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "Invalid token expiration",
			}
			c.JSON(http.StatusUnauthorized, errorResponse)
			c.Abort()
			return
		}

		if int64(exp) <= currentTime {
			log.Error().Err(err).Msg("Expired token")
			errorResponse := ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "Expired token",
			}
			c.JSON(http.StatusUnauthorized, errorResponse)
			c.Abort()
			return
		}
		wallet_address, ok := payload["wallet_address"].(string)
		if !ok {
			log.Error().Err(err).Msg("Invalid walletAddress")
			errorResponse := ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "Invalid walletAddress value",
			}
			c.JSON(http.StatusUnauthorized, errorResponse)
			c.Abort()
			return
		}

		log.Info().Msgf("getting admin using extracted walletAdress from token")
		_, err = repositories.GetAdminByWallet(serverConfigs.DB, wallet_address)
		if err != nil {
			log.Error().Err(err).Msgf("UnAuthorized: %s", err)
			errorResponse := ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "UnAuthorized",
			}
			c.JSON(http.StatusUnauthorized, errorResponse)
			c.Abort()
			return
		}

		c.Next()
	}
}

func IsAdminOrUserAuthenticated(serverConfigs utils.ServerConfigs) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info().Msg("auth middleware")

		authHeader := c.GetHeader("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")

		encryptionSecret, err := utils.GetDerivedEncryptionKey(serverConfigs.NEXTAUTH_SECRET, "")
		if err != nil {
			log.Error().Err(err).Msg("not authorized")
			errorResponse := ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "not authorized",
			}
			c.JSON(http.StatusUnauthorized, errorResponse)
			c.Abort()
			return
		}

		decryptedPayload, err := utils.DecodeAndDecryptToken(token, encryptionSecret)
		if err != nil {
			log.Error().Err(err).Msg("not authorized")
			errorResponse := ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "not authorized",
			}
			c.JSON(http.StatusUnauthorized, errorResponse)
			c.Abort()
			return
		}
		var payload map[string]interface{}
		err = json.Unmarshal([]byte(decryptedPayload), &payload)
		if err != nil {
			log.Error().Err(err).Msg("not authorized")
			errorResponse := ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "not authorized",
			}
			c.JSON(http.StatusUnauthorized, errorResponse)
			c.Abort()
			return
		}

		currentTime := time.Now().Unix()
		exp, ok := payload["exp"].(float64)
		if !ok {
			log.Error().Err(err).Msg("invalid or missing 'exp' field in payload")
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid or missing expiration in payload"})
			c.Abort()
			return
		}

		if int64(exp) <= currentTime {
			log.Error().Err(err).Msgf("Expired token")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Expired token"})
			c.Abort()
			return
		}

		userId, ok := payload["sub"].(string)
		if !ok {
			log.Error().Err(err).Msg("invalid or missing 'sub' field in payload")
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid or missing userID in payload"})
			c.Abort()
			return
		}

		var IsAdmin bool

		walletAddress, _ := payload["wallet_address"].(string)
		if walletAddress != "" {
			log.Info().Msgf("getting admin by walletAdress")
			admin, _ := repositories.GetAdminByWallet(serverConfigs.DB, walletAddress)
			if admin != nil {
				IsAdmin = true
			}
		}

		log.Info().Msgf("setting user and userId to the context for further access inside the handlers")
		if IsAdmin {
			c.Set("user", "admin")
			c.Set("userID", "userId")
		} else {
			c.Set("user", "user")
			c.Set("userID", userId)
		}

		c.Next()
	}
}
