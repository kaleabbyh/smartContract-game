package main-two

import (
	"encoding/json"
	"fmt"

	"crypto/sha256"

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

func decodeAndDecryptToken(encodedJWT string, encryptionSecret []byte) ([]byte, error) {
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

func main() {

	var encryptionSecret, err = GetDerivedEncryptionKey(NEXTAUTH_SECRET, "")
	encodedJWT := "eyJhbGciOiJkaXIiLCJlbmMiOiJBMjU2R0NNIn0..sYThWdGn2TTSa6SD.6ZVQQiSa6BVjgfIRUhKTuQDLTd5H96poHL4it986jqnNveyTpQ-pDvGZUKu0Eu7NgUd3FjU66FNj8uMH-3v9hyVe-leG4Hmin_YVmslOzQAOT-gQUL1KMFqT7FOOlPw47dbiVPhDDlhIl2wzTKxT1e7XsvZWgUp4RDZpd9u2ll6EKJQOCSEDB4OwuUam5TvqUgzWhyQnaeQxLTIk4iRxgLjFpb5rLRLEKn44P2yD6WvhPqg00cBFd2kJb93c21onAzsLziMA0QATb-P2wsIcSbD_UotpjUZKIkjbQiPvAJjEIIzpacT9FB35PrVTsLyMz57A83khOcLYrUhxOhSfnDXiE3RvDQC1hp9qqJ9U7I7RAO_Z2IA_YWdVHvriGiZpIffqbrQ5qtZ0m7dHCXZdRp9FQIPsSl8vjzonsCkXptvY0RyyyY0_xpFYV-GxgwAuWHenrOyCdQLEk9DIx3Vh1hGV4oxFEoZKNBfHd6oZI70IMQd7AJ-m9t_be0ge_C79svojN6v9ZJ1quINfl3gHK6WmGh8Qu_LjWBrKWwaDeRtz6JMcBhsrQSks1mwO0TBDj1z4Sg.nZucSw8tooCREUJ1az2_EQ"

	decryptedPayload, err := decodeAndDecryptToken(encodedJWT, encryptionSecret)
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
