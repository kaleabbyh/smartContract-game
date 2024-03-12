package advanced_encryption

import (
	"encoding/json"
	"fmt"
	"os"

	"crypto/sha256"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/square/go-jose"

	"golang.org/x/crypto/hkdf"
)

var NEXTAUTH_SECRET = "qa/nyfFrH+fXGH3dHMFpSDsZWbMs+Rk8aabOA4bGkwI="

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

func CheckToken() {
	var logger = log.Logger.With().Str("func", "Configs").Logger()

	if err := godotenv.Load(); err != nil {
		logger.Error().Err(err).Msg("Error loading .env file:")
	}
	AccessToken := os.Getenv("TOKEN")

	var encryptionSecret, err = GetDerivedEncryptionKey(NEXTAUTH_SECRET, "")
	encodedJWT := AccessToken

	decryptedPayload, err := DecodeAndDecryptToken(encodedJWT, encryptionSecret)
	if err != nil {
		fmt.Println("Error decoding and decrypting token:", err)
		return
	}

	var payload map[string]interface{}
	err = json.Unmarshal([]byte(decryptedPayload), &payload)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	name, ok := payload["name"].(string)
	if !ok {
		fmt.Println("Name field not found or not a string")
		return
	}

	fmt.Println("Name:", name)
	fmt.Println("Decoded and decrypted payload:", string(decryptedPayload))
}
