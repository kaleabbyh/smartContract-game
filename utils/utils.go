package utils

import (
	"fmt"

	"errors"
	"time"

	"crypto/sha256"

	"github.com/square/go-jose"
	"golang.org/x/crypto/hkdf"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type CustomJWTClaims struct {
	Role   string `json:"role"`
	UserId string `json:"userId"`

	jwt.StandardClaims
}

func HashPassword(password string) (string, error) {

	logger := log.Logger.With().Str("func", "HashPassword").Logger()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to hash password")
		return "", fmt.Errorf("could not hash password: %w", err)
	}

	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}

func GenerateToken(newClaims CustomJWTClaims, secret string) (string, error) {

	logger := log.Logger.With().Str("func", "GenerateToken").Logger()

	claims := &CustomJWTClaims{
		UserId: newClaims.UserId,
		Role:   newClaims.Role,

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		logger.Error().Err(err).Msg("could not generate token")
		return "", fmt.Errorf("could not generate token: %w", err)

	}

	return tokenString, nil
}

func VerifyToken(tokenString string, secret string) (*CustomJWTClaims, error) {

	logger := log.Logger.With().Str("func", "VerifyToken").Logger()

	token, err := jwt.ParseWithClaims(tokenString, &CustomJWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

	if err != nil {
		logger.Error().Err(err).Msg("could not validate token")
		return nil, fmt.Errorf("could not validate token: %w", err)
	}

	if claims, ok := token.Claims.(*CustomJWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func GetDerivedEncryptionKey(keyMaterial, salt string) ([]byte, error) {
	hkdf := hkdf.New(sha256.New, []byte(keyMaterial), []byte(salt), []byte(fmt.Sprintf("NextAuth.js Generated Encryption Key%s", salt)))
	key := make([]byte, 32)
	_, err := hkdf.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func DecodeAndDecryptToken(encodedJWT string, encryptionSecret []byte) ([]byte, error) {
	object, err := jose.ParseEncrypted(encodedJWT)
	if err != nil {
		return nil, err
	}

	payload, err := object.Decrypt(encryptionSecret)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

func GenerateUUID() (string, error) {
	id, err := uuid.New().MarshalText()
	if err != nil {

		return "", err
	}
	ID := string(id)
	return ID, nil
}
